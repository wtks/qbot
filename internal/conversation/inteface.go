package conversation

import (
	"github.com/wtks/qbot/pkg/logging"
	"github.com/wtks/qbot/pkg/qapi"
)

type QBotInterface interface {
	AddMessageHook(mid string, ctx *Context)
	RemoveMessageHook(ctx *Context)
	FinishContext(ctx *Context)
	GetAPIClient() *qapi.WrappedClient
	GetLogger() logging.Logger
}
