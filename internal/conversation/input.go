package conversation

import (
	"github.com/wtks/qbot/pkg/conversation"
	"github.com/wtks/qbot/pkg/input"
)

type timeoutInput struct{}

type resultInput struct{}

type errorInput struct {
	Error error
}

type inputImpl struct {
	previousStep   int
	previousResult interface{}
	input          interface{}
}

func (i *inputImpl) PreviousStep() int {
	return i.previousStep
}

func (i *inputImpl) PreviousResult() interface{} {
	return i.previousResult
}

func (i *inputImpl) Type() int {
	switch i.input.(type) {
	case input.Message:
		return conversation.InputTypeMessage
	case input.MessageStamp:
		return conversation.InputTypeStamp
	case timeoutInput:
		return conversation.InputTypeTimeout
	case resultInput:
		return conversation.InputTypeResult
	case errorInput:
		return conversation.InputTypeError
	default:
		panic("unknown type")
	}
}

func (i *inputImpl) M() input.Message {
	m, _ := i.input.(input.Message)
	return m
}

func (i *inputImpl) S() input.MessageStamp {
	s, _ := i.input.(input.MessageStamp)
	return s
}

func (i *inputImpl) E() error {
	e, _ := i.input.(errorInput)
	return e.Error
}
