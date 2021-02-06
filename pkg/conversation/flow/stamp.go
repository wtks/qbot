package flow

import (
	"github.com/wtks/qbot/pkg/conversation"
	"github.com/wtks/qbot/pkg/conversation/cond"
)

func Stamp(names ...string) conversation.FlowBuilder {
	return conversation.Case(cond.Stamp(names...))
}

func StampID(ids ...string) conversation.FlowBuilder {
	return conversation.Case(cond.StampID(ids...))
}

func StampC(key string) conversation.FlowBuilder {
	return conversation.Case(cond.CStamp(key))
}

func AnyStamp() conversation.FlowBuilder {
	return conversation.Case(cond.AnyStamp())
}
