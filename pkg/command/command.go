package command

type Trigger func(input Input) bool

type ExecuteFunc func(ctx Context, input Input) error

type Command interface {
	GetName() string
	Trigger(input Input) bool
	Execute(ctx Context, input Input) error
}
