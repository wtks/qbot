package conversation

import (
	"fmt"
	"github.com/wtks/qbot/internal/input"
	"github.com/wtks/qbot/pkg/conversation"
	"github.com/wtks/qbot/pkg/event"
	"sync"
	"time"
)

type Context struct {
	key            string
	dm             bool
	channelID      string
	userID         string
	spec           *SpecImpl
	currentStep    *StepImpl
	previousResult interface{}
	values         map[string]interface{}
	running        bool
	timeoutTimer   *time.Timer
	statusLock     sync.RWMutex

	qbot QBotInterface
}

func (c *Context) IsDM() bool {
	return c.dm
}

func (c *Context) GetChannelID() string {
	return c.channelID
}

func (c *Context) GetUserID() string {
	return c.userID
}

func (c *Context) GetUserName() string {
	panic("implement me")
}

func (c *Context) Get(key string) (interface{}, bool) {
	v, ok := c.values[key]
	return v, ok
}

func (c *Context) MustGet(key string) interface{} {
	v, ok := c.values[key]
	if !ok {
		panic(fmt.Errorf("the Context didn't have the key's value: %s", key))
	}
	return v
}

func (c *Context) Put(key string, value interface{}) {
	c.values[key] = value
}

func (c *Context) Delete(key string) {
	delete(c.values, key)
}

func (c *Context) GetCurrentStep() *StepImpl {
	c.statusLock.RLock()
	defer c.statusLock.RUnlock()
	return c.currentStep
}

func (c *Context) Execute(step *StepImpl, in interface{}) {
	c.statusLock.Lock()
	if c.running {
		c.statusLock.Unlock()
		return
	}
	c.running = true

	// timeoutタイマー停止
	if c.timeoutTimer != nil {
		c.timeoutTimer.Stop()
		c.timeoutTimer = nil
	}

	// messageHook削除
	if c.currentStep != nil && len(c.currentStep.stampMatchers) > 0 {
		c.qbot.RemoveMessageHook(c)
	}

	prev := c.currentStep
	prevResult := c.previousResult
	c.currentStep = step
	c.previousResult = nil
	c.statusLock.Unlock()

	var prevID int
	if prev == nil {
		prevID = conversation.StepStart
	} else {
		prevID = prev.id
	}

	res := c.currentStep.f(c, &inputImpl{
		previousStep:   prevID,
		previousResult: prevResult,
		input:          in,
	})
	if res == nil {
		c.qbot.FinishContext(c)

		c.statusLock.Lock()
		c.running = false
		c.statusLock.Unlock()
		return
	}

	if err := res.StepError; err != nil {
		c.statusLock.Lock()
		c.running = false
		c.statusLock.Unlock()

		c.Execute(c.spec.ErrorStep, errorInput{Error: err})
		return
	} else {
		var responseMID string
		if res.ResponseMessage != nil {
			// メッセージ投稿
			responseMID, err = c.qbot.GetAPIClient().PostMessage(c.GetChannelID(), *res.ResponseMessage)
			if err != nil {
				c.qbot.GetLogger().Error(fmt.Errorf("failed to PostMessage (cid:%s): %w", c.GetChannelID(), err))
			}
			for _, stamp := range res.ResponseMessageStamps {
				if err := c.qbot.GetAPIClient().AddMessageStamp(responseMID, stamp); err != nil {
					c.qbot.GetLogger().Error(fmt.Errorf("failed to AddMessageStamp (mid:%s, stamp:%s): %w", responseMID, stamp, err))
				}
			}
		}
		if len(res.ResponseStamp) > 0 {
			// スタンプ押す
			if m, ok := in.(*input.Message); ok {
				for _, s := range res.ResponseStamp {
					if err := c.qbot.GetAPIClient().AddMessageStamp(m.MessageID(), s); err != nil {
						c.qbot.GetLogger().Error(fmt.Errorf("failed to AddMessageStamp (mid:%s, stamp:%s): %w", m.MessageID(), s, err))
					}
				}
			}
		}

		c.previousResult = res.StepResult
		if next, ok := c.currentStep.ResultMatching(c, res.StepResult); ok {
			c.statusLock.Lock()
			c.running = false
			c.statusLock.Unlock()

			c.Execute(next, resultInput{})
			return
		}

		if len(c.currentStep.stampMatchers) > 0 && len(responseMID) > 0 {
			c.qbot.AddMessageHook(responseMID, c)
		}

		if c.currentStep.timeoutFlow != nil {
			to := c.currentStep.timeoutFlow.To
			c.timeoutTimer = time.AfterFunc(c.currentStep.timeoutFlow.Duration, func() {
				c.Execute(to, timeoutInput{})
			})
		}
	}

	if c.currentStep.end {
		c.qbot.FinishContext(c)
	}

	c.statusLock.Lock()
	c.running = false
	c.statusLock.Unlock()
}

func (c *Context) GetContextKey() string {
	return c.key
}

func GetContextKey(p event.Message) string {
	return p.ChannelID + ":" + p.User.ID
}

func NewContext(qbot QBotInterface, key string, spec *SpecImpl, p event.Message, isDM bool) *Context {
	return &Context{
		key:       key,
		dm:        isDM,
		channelID: p.ChannelID,
		userID:    p.User.ID,
		spec:      spec,
		values:    map[string]interface{}{},
		qbot:      qbot,
	}
}
