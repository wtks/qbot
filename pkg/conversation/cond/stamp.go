package cond

import (
	"github.com/wtks/qbot/pkg/conversation"
	"github.com/wtks/qbot/pkg/input"
)

func Stamp(names ...string) conversation.StampMatcher {
	return func(_ conversation.Context, input input.MessageStamp) bool {
		for _, name := range names {
			if name == input.StampName() {
				return true
			}
		}
		return false
	}
}

func StampID(ids ...string) conversation.StampMatcher {
	return func(_ conversation.Context, input input.MessageStamp) bool {
		for _, id := range ids {
			if id == input.StampID() {
				return true
			}
		}
		return false
	}
}

func CStamp(key string) conversation.StampMatcher {
	return func(ctx conversation.Context, input input.MessageStamp) bool {
		v, ok := ctx.Get(key)
		if !ok {
			return false
		}
		switch s := v.(type) {
		case string:
			return input.StampName() == s
		case []string:
			for _, name := range s {
				if name == input.StampName() {
					return true
				}
			}
		}
		return false
	}
}

func AnyStamp() conversation.StampMatcher {
	return func(_ conversation.Context, _ input.MessageStamp) bool {
		return true
	}
}
