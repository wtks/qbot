package conversation

import "github.com/wtks/qbot/pkg/input"

type Context interface {
	IsDM() bool
	GetChannelID() string
	GetUserID() string
	GetUserName() string

	Get(key string) (interface{}, bool)
	MustGet(key string) interface{}
	Put(key string, value interface{})
	Delete(key string)
}

type TriggerMatcher func(input input.Message) bool

type Spec struct {
	Name         string
	Trigger      TriggerMatcher
	Steps        []Step
	AllowDM      bool
	AllowChannel bool
}

type StepFunc func(ctx Context, input Input) *Output

type Step struct {
	Name  int
	Func  StepFunc
	Next  []Flow
	Start bool
}
