package input

import (
	"github.com/wtks/qbot/pkg/event"
	"time"
)

type Message struct {
	p    event.Message
	isDM bool
}

func (m *Message) MessageID() string {
	return m.p.ID
}

func (m *Message) UserID() string {
	return m.p.User.ID
}

func (m *Message) ChannelID() string {
	return m.p.ChannelID
}

func (m *Message) Message() string {
	return m.p.PlainText
}

func (m *Message) SentAt() time.Time {
	return m.p.CreatedAt
}

func (m *Message) IsDM() bool {
	return m.isDM
}

func (m *Message) RawPayload() event.Message {
	return m.p
}

func ConvertFromMessagePayload(p event.Message, isDM bool) *Message {
	return &Message{p: p, isDM: isDM}
}
