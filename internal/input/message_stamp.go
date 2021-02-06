package input

import (
	"github.com/wtks/qbot/pkg/event"
	"github.com/wtks/qbot/pkg/qapi"
	"time"
)

type MessageStamp struct {
	p event.MessageStamp
	c *qapi.WrappedClient
}

func (m *MessageStamp) UserID() string {
	return m.p.UserID
}

func (m *MessageStamp) UserName() string {
	return m.c.GetUserNameByID(m.p.UserID)
}

func (m *MessageStamp) StampID() string {
	return m.p.StampID
}

func (m *MessageStamp) StampName() string {
	return m.c.GetStampNameByID(m.p.StampID)
}

func (m *MessageStamp) CreatedAt() time.Time {
	return m.p.CreatedAt
}

func ConvertFromMessageStampPayload(p event.MessageStamp, c *qapi.WrappedClient) *MessageStamp {
	return &MessageStamp{p: p, c: c}
}
