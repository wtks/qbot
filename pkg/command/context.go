package command

type Context interface {
	Command() Command
	Reply(message string) error
	ReplyViaDM(message string) error
	PushStamp(stamp string) error

	Get(key string) (interface{}, bool)
	MustGet(key string) interface{}
	Put(key string, value interface{})
	Delete(key string)
}
