package qbot

import (
	"fmt"
	"github.com/asaskevich/EventBus"
	"github.com/robfig/cron/v3"
	internalconv "github.com/wtks/qbot/internal/conversation"
	"github.com/wtks/qbot/internal/header"
	"github.com/wtks/qbot/internal/httputils"
	"github.com/wtks/qbot/internal/input"
	"github.com/wtks/qbot/internal/utils"
	"github.com/wtks/qbot/pkg/conversation"
	"github.com/wtks/qbot/pkg/event"
	"github.com/wtks/qbot/pkg/logging"
	"github.com/wtks/qbot/pkg/qapi"
	"net/http"
	"sync"
)

type QBot struct {
	verificationToken string
	apiClient         *qapi.WrappedClient
	internal          *internalInterface

	cron *cron.Cron

	eventChan chan interface{}
	eventbus  EventBus.Bus

	conversationSpecs []*internalconv.SpecImpl
	specsLock         sync.RWMutex

	conversationContexts     map[string]*internalconv.Context
	hookMessages             map[string]*internalconv.Context
	conversationContextsLock sync.Mutex

	logger logging.Logger
}

func New(verificationToken, accessToken string) (*QBot, error) {
	wc, err := qapi.NewWrappedClient(accessToken)
	if err != nil {
		return nil, err
	}
	qbot := &QBot{
		verificationToken:    verificationToken,
		apiClient:            wc,
		eventChan:            make(chan interface{}, 100),
		eventbus:             EventBus.New(),
		conversationSpecs:    make([]*internalconv.SpecImpl, 0),
		conversationContexts: map[string]*internalconv.Context{},
		hookMessages:         map[string]*internalconv.Context{},
		internal:             &internalInterface{},
		logger:               &logging.StdLogger{},
	}
	qbot.internal.QB = qbot

	go qbot.loop()
	return qbot, nil
}

func (QB *QBot) AddCommand() {
	QB.specsLock.Lock()
	defer QB.specsLock.Unlock()
	// TODO command
}

func (QB *QBot) AddConversation(conversation conversation.Spec) error {
	spec, err := internalconv.BuildSpec(conversation)
	if err != nil {
		return err
	}
	QB.specsLock.Lock()
	defer QB.specsLock.Unlock()
	QB.conversationSpecs = append(QB.conversationSpecs, spec)
	return nil
}

func (QB *QBot) AddCron(spec string, f func(qb *QBot)) error {
	QB.specsLock.Lock()
	defer QB.specsLock.Unlock()
	_, err := QB.cron.AddFunc(spec, func() {
		f(QB)
	})
	return err
}

func (QB *QBot) AddHandlers(handler interface{}) error {
	t, ok := event.GetEventFuncType(handler)
	if !ok {
		return fmt.Errorf("invalid handler func")
	}
	return QB.eventbus.SubscribeAsync(t, handler, false)
}

func (QB *QBot) RemoveHandlers(handler interface{}) error {
	t, ok := event.GetEventFuncType(handler)
	if !ok {
		return fmt.Errorf("invalid handler func")
	}
	return QB.eventbus.Unsubscribe(t, handler)
}

func (QB *QBot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// トークン検証
	if r.Header.Get(header.BotToken) != QB.verificationToken {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	// Json POSTのみ受付
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	if !httputils.IsJsonRequest(r) {
		http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
		return
	}

	ev := r.Header.Get(header.BotEvent)
	switch ev {
	case event.Ping:
		var payload event.PingPayload
		if err := httputils.DecodeJSON(r, &payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		QB.eventbus.Publish(ev, payload)
	case event.Left:
		var payload event.LeftPayload
		if err := httputils.DecodeJSON(r, &payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		QB.eventbus.Publish(ev, payload)
	case event.Joined:
		var payload event.JoinedPayload
		if err := httputils.DecodeJSON(r, &payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		QB.eventbus.Publish(ev, payload)
	case event.MessageCreated:
		var payload event.MessageCreatedPayload
		if err := httputils.DecodeJSON(r, &payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		QB.eventbus.Publish(ev, payload)
		QB.eventChan <- &payload
	case event.MessageDeleted:
		var payload event.MessageDeletedPayload
		if err := httputils.DecodeJSON(r, &payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		QB.eventbus.Publish(ev, payload)
	case event.MessageUpdated:
		var payload event.MessageUpdatedPayload
		if err := httputils.DecodeJSON(r, &payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		QB.eventbus.Publish(ev, payload)
	case event.BotMessageStampsUpdated:
		var payload event.BotMessageStampsUpdatedPayload
		if err := httputils.DecodeJSON(r, &payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		QB.eventbus.Publish(ev, payload)
		QB.eventChan <- &payload
	case event.DirectMessageCreated:
		var payload event.DirectMessageCreatedPayload
		if err := httputils.DecodeJSON(r, &payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		QB.eventbus.Publish(ev, payload)
		QB.eventChan <- &payload
	case event.DirectMessageDeleted:
		var payload event.DirectMessageDeletedPayload
		if err := httputils.DecodeJSON(r, &payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		QB.eventbus.Publish(ev, payload)
	case event.DirectMessageUpdated:
		var payload event.DirectMessageUpdatedPayload
		if err := httputils.DecodeJSON(r, &payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		QB.eventbus.Publish(ev, payload)
	case event.ChannelCreated:
		var payload event.ChannelCreatedPayload
		if err := httputils.DecodeJSON(r, &payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		QB.eventbus.Publish(ev, payload)
	case event.ChannelTopicChanged:
		var payload event.ChannelTopicChangedPayload
		if err := httputils.DecodeJSON(r, &payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		QB.eventbus.Publish(ev, payload)
	case event.UserCreated:
		var payload event.UserCreatedPayload
		if err := httputils.DecodeJSON(r, &payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		QB.eventbus.Publish(ev, payload)
	case event.StampCreated:
		var payload event.StampCreatedPayload
		if err := httputils.DecodeJSON(r, &payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		QB.eventbus.Publish(ev, payload)
	case event.TagAdded:
		var payload event.TagAddedPayload
		if err := httputils.DecodeJSON(r, &payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		QB.eventbus.Publish(ev, payload)
	case event.TagRemoved:
		var payload event.TagRemovedPayload
		if err := httputils.DecodeJSON(r, &payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		QB.eventbus.Publish(ev, payload)
	default:
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}

func (QB *QBot) loop() {
	for {
		select {
		case ev := <-QB.eventChan:
			switch payload := ev.(type) {
			case *event.BotMessageStampsUpdatedPayload:
				QB.conversationContextsLock.Lock()
				if ctx, ok := QB.hookMessages[payload.MessageID]; ok {
					// 会話コンテキストが存在
					if s, ok := utils.ExtractFirstStamp(payload.Stamps, ctx.GetUserID()); ok {
						inputStamp := input.ConvertFromMessageStampPayload(s, QB.apiClient)
						next, ok := ctx.GetCurrentStep().StampMatching(ctx, inputStamp)
						if ok {
							// 次のステップを実行
							delete(QB.hookMessages, payload.MessageID)
							go ctx.Execute(next, inputStamp)
						}
					}
				}
				QB.conversationContextsLock.Unlock()
			case *event.DirectMessageCreatedPayload:
				QB.processMessageCreated(payload.Message, true)
			case *event.MessageCreatedPayload:
				QB.processMessageCreated(payload.Message, false)
			}
		}
	}
}

func (QB *QBot) processMessageCreated(message event.Message, isDM bool) {
	inputMessage := input.ConvertFromMessagePayload(message, isDM)

	// 会話コンテキストをチェック
	convctx, ok := QB.getConvContext(internalconv.GetContextKey(message))
	if ok {
		// 会話コンテキストが存在
		next, ok := convctx.GetCurrentStep().TextMatching(convctx, inputMessage)
		if ok {
			// 次のステップを実行
			go convctx.Execute(next, inputMessage)
			return
		}
	}

	// コマンド仕様を探索

	// 会話仕様を探索
	QB.specsLock.RLock()
	var convSpec *internalconv.SpecImpl
	for _, spec := range QB.conversationSpecs {
		if spec.Trigger(inputMessage) {
			// 会話spec発見
			convSpec = spec
			break
		}
	}
	QB.specsLock.RUnlock()
	if convSpec != nil {
		// 会話開始
		convctx = QB.newConvContext(convSpec, message, isDM)
		// ステップ実行
		go convctx.Execute(convSpec.StartStep, inputMessage)
		return
	}
}

func (QB *QBot) getConvContext(key string) (*internalconv.Context, bool) {
	QB.conversationContextsLock.Lock()
	defer QB.conversationContextsLock.Unlock()
	ctx, ok := QB.conversationContexts[key]
	return ctx, ok

}

func (QB *QBot) newConvContext(spec *internalconv.SpecImpl, m event.Message, isDM bool) *internalconv.Context {
	ctx := internalconv.NewContext(QB.internal, internalconv.GetContextKey(m), spec, m, isDM)

	QB.conversationContextsLock.Lock()
	defer QB.conversationContextsLock.Unlock()
	QB.conversationContexts[ctx.GetContextKey()] = ctx
	return ctx
}

type internalInterface struct {
	QB *QBot
}

func (i *internalInterface) AddMessageHook(mid string, ctx *internalconv.Context) {
	i.QB.conversationContextsLock.Lock()
	defer i.QB.conversationContextsLock.Unlock()
	i.QB.hookMessages[mid] = ctx
}

func (i *internalInterface) RemoveMessageHook(ctx *internalconv.Context) {
	i.QB.conversationContextsLock.Lock()
	defer i.QB.conversationContextsLock.Unlock()
	for key, _ctx := range i.QB.hookMessages {
		if ctx == _ctx {
			delete(i.QB.hookMessages, key)
			return
		}
	}
}

func (i *internalInterface) FinishContext(ctx *internalconv.Context) {
	i.QB.conversationContextsLock.Lock()
	defer i.QB.conversationContextsLock.Unlock()
	delete(i.QB.conversationContexts, ctx.GetContextKey())
}

func (i *internalInterface) GetAPIClient() *qapi.WrappedClient {
	return i.QB.apiClient
}

func (i *internalInterface) GetLogger() logging.Logger {
	return i.QB.logger
}
