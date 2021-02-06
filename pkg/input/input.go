package input

import (
	"github.com/wtks/qbot/pkg/event"
	"time"
)

type Message interface {
	MessageID() string
	UserID() string
	ChannelID() string
	Message() string
	SentAt() time.Time
	IsDM() bool
	RawPayload() event.Message
}

type MessageStamp interface {
	UserID() string
	UserName() string
	StampID() string
	StampName() string
	CreatedAt() time.Time
}
