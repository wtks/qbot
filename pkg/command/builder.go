package command

type simpleCommand struct {
	name    string
	trigger Trigger
	f       ExecuteFunc
}

func NewSimple(name string, trigger Trigger, f ExecuteFunc) Command {
	return &simpleCommand{
		name:    name,
		trigger: trigger,
		f:       f,
	}
}

func (s *simpleCommand) GetName() string {
	return s.name
}

func (s *simpleCommand) Trigger(input Input) bool {
	return s.trigger(input)
}

func (s *simpleCommand) Execute(ctx Context, input Input) error {
	return s.f(ctx, input)
}
