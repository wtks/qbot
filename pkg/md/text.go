package md

import (
	"fmt"
	"strings"
)

func Link(desc, url string) string {
	return fmt.Sprintf("[%s](%s)", desc, url)
}

func (m *MD) WriteLink(desc, url string) *MD {
	m.Write(Link(desc, url))
	return m
}

func (m *MD) WriteLinkLn(desc, url string) *MD {
	m.WriteLink(desc, url)
	m.Writeln()
	return m
}

func Code(text string) string {
	return fmt.Sprintf("`%s`", text)
}

func (m *MD) WriteCode(text string) *MD {
	m.Write(Code(text))
	return m
}

func (m *MD) WriteCodeLn(text string) *MD {
	m.WriteCode(text)
	m.Writeln()
	return m
}

func MultiCode(lines string, contentType ...string) string {
	if len(contentType) > 0 {
		return fmt.Sprintf("```%s\n%s\n```\n", contentType[0], lines)
	}
	return fmt.Sprintf("```\n%s\n```\n", lines)
}

func (m *MD) WriteMultiCode(lines string, contentType ...string) *MD {
	m.Write(MultiCode(lines, contentType...))
	return m
}

func Italic(text string) string {
	return fmt.Sprintf("*%s*", text)
}

func (m *MD) WriteItalic(text string) *MD {
	m.Write(Italic(text))
	return m
}

func (m *MD) WriteItalicLn(text string) *MD {
	m.WriteItalic(text)
	m.Writeln()
	return m
}

func Bold(text string) string {
	return fmt.Sprintf("**%s**", text)
}

func (m *MD) WriteBold(text string) *MD {
	m.Write(Bold(text))
	return m
}

func (m *MD) WriteBoldLn(text string) *MD {
	m.WriteBold(text)
	m.Writeln()
	return m
}

func Strike(text string) string {
	return fmt.Sprintf("~~%s~~", text)
}

func (m *MD) WriteStrike(text string) *MD {
	m.Write(Strike(text))
	return m
}

func (m *MD) WriteStrikeLn(text string) *MD {
	m.WriteStrike(text)
	m.Writeln()
	return m
}

func BlockQuotes(text string) string {
	var b strings.Builder
	for _, line := range strings.Split(text, "\n") {
		b.WriteString("> ")
		b.WriteString(line)
		b.WriteString("\n")
	}
	return b.String()
}

func (m *MD) WriteBlockQuotes(text string) *MD {
	m.Write(BlockQuotes(text))
	return m
}
