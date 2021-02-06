package cond

import (
	"github.com/wtks/qbot/pkg/conversation"
	"github.com/wtks/qbot/pkg/input"
	"regexp"
	"strconv"
	"strings"
)

func Regex(pattern string) conversation.TextMatcher {
	r := regexp.MustCompile(pattern)
	return func(ctx conversation.Context, input input.Message) bool {
		return r.MatchString(input.Message())
	}
}

func Text(s string) conversation.TextMatcher {
	return func(ctx conversation.Context, input input.Message) bool {
		return strings.TrimSpace(input.Message()) == s
	}
}

func IntText() conversation.TextMatcher {
	return func(ctx conversation.Context, input input.Message) bool {
		m := strings.TrimSpace(input.Message())
		_, err := strconv.ParseInt(m, 10, 64)
		return err == nil
	}
}

func FloatText() conversation.TextMatcher {
	return func(ctx conversation.Context, input input.Message) bool {
		m := strings.TrimSpace(input.Message())
		_, err := strconv.ParseFloat(m, 64)
		return err == nil
	}
}
