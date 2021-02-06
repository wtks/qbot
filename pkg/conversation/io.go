package conversation

import (
	"github.com/wtks/qbot/pkg/input"
)

const (
	InputTypeMessage = iota
	InputTypeStamp
	InputTypeTimeout
	InputTypeResult
	InputTypeError
)

type Input interface {
	PreviousStep() int
	PreviousResult() interface{}
	Type() int
	M() input.Message
	S() input.MessageStamp
	E() error
}

type Output struct {
	StepResult            interface{}
	StepError             error
	ResponseMessage       *string
	ResponseMessageStamps []string
	ResponseStamp         []string
}

func NewOutput() *Output {
	return &Output{}
}

func (o *Output) Result(result interface{}) *Output {
	o.StepResult = result
	return o
}

func (o *Output) Message(message string, stamps ...string) *Output {
	o.ResponseMessage = &message
	o.ResponseMessageStamps = stamps
	return o
}

func (o *Output) Stamps(stamps ...string) *Output {
	o.ResponseStamp = append(o.ResponseStamp, stamps...)
	return o
}

func (o *Output) Error(err error) *Output {
	o.StepError = err
	return o
}

func Error(err error) *Output {
	return NewOutput().Error(err)
}
