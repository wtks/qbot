package conversation

import (
	"github.com/wtks/qbot/pkg/input"
	"time"
)

type FlowType int

const (
	FlowTypeText FlowType = iota
	FlowTypeStamp
	FlowTypeTimeout
	FlowTypeResult
)

type FlowMatcher interface {
	GetEdgeType() FlowType
}

type Flow struct {
	Match FlowMatcher
	To    int
}

type TimeoutMatcher time.Duration

func (c TimeoutMatcher) GetEdgeType() FlowType {
	return FlowTypeTimeout
}

type TextMatcher func(ctx Context, input input.Message) bool

func (c TextMatcher) GetEdgeType() FlowType {
	return FlowTypeText
}

type StampMatcher func(ctx Context, input input.MessageStamp) bool

func (c StampMatcher) GetEdgeType() FlowType {
	return FlowTypeStamp
}

type ResultMatcher func(ctx Context, result interface{}) bool

func (c ResultMatcher) GetEdgeType() FlowType {
	return FlowTypeResult
}

type FlowBuilder struct {
	match FlowMatcher
}

func Case(cond FlowMatcher) FlowBuilder {
	return FlowBuilder{match: cond}
}

func (f FlowBuilder) GoTo(step int) Flow {
	return Flow{Match: f.match, To: step}
}
