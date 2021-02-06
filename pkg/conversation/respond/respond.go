package respond

import "github.com/wtks/qbot/pkg/conversation"

func Message(message string, stamps ...string) *conversation.Output {
	return conversation.NewOutput().Message(message, stamps...)
}

func Stamps(stamps ...string) *conversation.Output {
	return conversation.NewOutput().Stamps(stamps...)
}

func Result(result interface{}) *conversation.Output {
	return conversation.NewOutput().Result(result)
}
