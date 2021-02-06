package flow

import (
	"github.com/wtks/qbot/pkg/conversation"
	"github.com/wtks/qbot/pkg/conversation/cond"
)

func TextPattern(pattern string) conversation.FlowBuilder {
	return conversation.Case(cond.Regex(pattern))
}

func Text(s string) conversation.FlowBuilder {
	return conversation.Case(cond.Text(s))
}

func AnyText() conversation.FlowBuilder {
	return conversation.Case(cond.Regex(".+"))
}

func IntText() conversation.FlowBuilder {
	return conversation.Case(cond.IntText())
}

func FloatText() conversation.FlowBuilder {
	return conversation.Case(cond.FloatText())
}
