package conversation

import (
	"github.com/wtks/qbot/pkg/conversation"
	"github.com/wtks/qbot/pkg/input"
	"time"
)

type SpecImpl struct {
	Name         string
	Trigger      conversation.TriggerMatcher
	Steps        map[int]*StepImpl
	StartStep    *StepImpl
	ErrorStep    *StepImpl
	AllowDM      bool
	AllowChannel bool
}

type StepImpl struct {
	id             int
	f              conversation.StepFunc
	stampMatchers  []*stampMatcher
	textMatchers   []*textMatcher
	resultMatchers []*resultMatcher
	timeoutFlow    *timeoutFlow

	end bool
}

func (step *StepImpl) StampMatching(ctx *Context, input input.MessageStamp) (*StepImpl, bool) {
	for _, matcher := range step.stampMatchers {
		if matcher.Match(ctx, input) {
			return matcher.To, true
		}
	}
	return nil, false
}

func (step *StepImpl) TextMatching(ctx *Context, input input.Message) (*StepImpl, bool) {
	for _, matcher := range step.textMatchers {
		if matcher.Match(ctx, input) {
			return matcher.To, true
		}
	}
	return nil, false
}

func (step *StepImpl) ResultMatching(ctx *Context, result interface{}) (*StepImpl, bool) {
	for _, matcher := range step.resultMatchers {
		if matcher.Match(ctx, result) {
			return matcher.To, true
		}
	}
	return nil, false
}

type stampMatcher struct {
	Match conversation.StampMatcher
	To    *StepImpl
}

type textMatcher struct {
	Match conversation.TextMatcher
	To    *StepImpl
}

type timeoutFlow struct {
	Duration time.Duration
	To       *StepImpl
}

type resultMatcher struct {
	Match conversation.ResultMatcher
	To    *StepImpl
}
