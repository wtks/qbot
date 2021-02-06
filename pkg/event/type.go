package event

import (
	"reflect"
)

const ANY = "___ANY___"

type AnyFunc func(payload interface{})
type AnyFuncInterface interface {
	Handle(payload interface{})
}

func GetEventFuncType(f interface{}) (string, bool) {
	t := reflect.TypeOf(f)
	is := func(t reflect.Type, f interface{}) bool {
		return t.ConvertibleTo(reflect.TypeOf(f))
	}

	switch {
	case is(t, TagAddedFunc(nil)):
		return TagAdded, true
	case is(t, TagRemovedFunc(nil)):
		return TagRemoved, true
	case is(t, PingFunc(nil)):
		return Ping, true
	case is(t, LeftFunc(nil)):
		return Left, true
	case is(t, JoinedFunc(nil)):
		return Joined, true
	case is(t, UserCreatedFunc(nil)):
		return UserCreated, true
	case is(t, StampCreatedFunc(nil)):
		return StampCreated, true
	case is(t, ChannelCreatedFunc(nil)):
		return ChannelCreated, true
	case is(t, ChannelTopicChangedFunc(nil)):
		return ChannelTopicChanged, true
	case is(t, MessageCreatedFunc(nil)):
		return MessageCreated, true
	case is(t, MessageUpdatedFunc(nil)):
		return MessageUpdated, true
	case is(t, MessageDeletedFunc(nil)):
		return MessageDeleted, true
	case is(t, DirectMessageCreatedFunc(nil)):
		return DirectMessageCreated, true
	case is(t, DirectMessageUpdatedFunc(nil)):
		return DirectMessageUpdated, true
	case is(t, DirectMessageDeletedFunc(nil)):
		return DirectMessageDeleted, true
	case is(t, BotMessageStampsUpdatedFunc(nil)):
		return BotMessageStampsUpdated, true
	case is(t, AnyFunc(nil)):
		return ANY, true
	default:
		return "", false
	}
}

func GetImplementedEventHandlerTypes(h interface{}) []string {
	var result []string

	if _, ok := h.(PingFuncInterface); ok {
		result = append(result, Ping)
	}
}
