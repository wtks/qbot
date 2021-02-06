package flow

import (
	"github.com/wtks/qbot/pkg/conversation"
	"github.com/wtks/qbot/pkg/conversation/cond"
	"time"
)

func Result(v interface{}) conversation.FlowBuilder {
	return conversation.Case(cond.Result(v))
}

func Immediately() conversation.FlowBuilder {
	return conversation.Case(cond.Immediately())
}

func Timeout(duration time.Duration) conversation.FlowBuilder {
	return conversation.Case(cond.Timeout(duration))
}
