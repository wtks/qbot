package md

import (
	"fmt"
	"strings"
)

type MD struct {
	builder strings.Builder
}

func New() *MD {
	return &MD{}
}

func (m *MD) Write(content ...string) *MD {
	if len(content) > 0 {
		m.builder.WriteString(strings.Join(content, " "))
	}
	return m
}

func (m *MD) Writeln(content ...string) *MD {
	m.Write(content...)
	m.Write("\n")
	return m
}

func (m *MD) Writef(format string, a ...interface{}) *MD {
	m.Write(fmt.Sprintf(format, a...))
	return m
}

func (m *MD) String() string {
	return m.builder.String()
}
