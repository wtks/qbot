package utils

import (
	"github.com/wtks/qbot/pkg/event"
	"time"
)

func ExtractFirstStamp(stamps []event.MessageStamp, userID string) (event.MessageStamp, bool) {
	i := -1
	t := time.Now().AddDate(1, 0, 0)
	for idx, s := range stamps {
		if s.UserID == userID {
			if t.After(s.CreatedAt) {
				t = s.CreatedAt
				i = idx
			}
		}
	}
	if i == -1 {
		return event.MessageStamp{}, false
	}
	return stamps[i], true
}
