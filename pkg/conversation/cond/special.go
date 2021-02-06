package cond

import (
	"github.com/wtks/qbot/pkg/conversation"
	"time"
)

func Result(v interface{}) conversation.ResultMatcher {
	return func(_ conversation.Context, result interface{}) bool {
		return result == v
	}
}

func Immediately() conversation.ResultMatcher {
	return func(_ conversation.Context, _ interface{}) bool {
		return true
	}
}

func Timeout(duration time.Duration) conversation.TimeoutMatcher {
	return conversation.TimeoutMatcher(duration)
}
