package trigger

type Func func(ctx Context, input Input) bool

type Context interface {
	Get(key string) (interface{}, bool)
	Put(key string, value interface{})
}

type Input interface {
}
