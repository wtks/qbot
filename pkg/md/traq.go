package md

import (
	"fmt"
	"github.com/wtks/qbot/pkg/qapi"
)

func Stamp(name string) string {
	return fmt.Sprintf(":%s:", name)
}

func (m *MD) WriteStamp(name string) *MD {
	m.Write(Stamp(name))
	return m
}

func (m *MD) WriteStampLn(name string) *MD {
	m.WriteStrike(name)
	m.Writeln()
	return m
}

func Tex(content string) string {
	return fmt.Sprintf("$%s$", content)
}

func (m *MD) WriteTex(content string) *MD {
	m.Write(Tex(content))
	return m
}

func (m *MD) WriteTexLn(content string) *MD {
	m.WriteTex(content)
	m.Writeln()
	return m
}

func TexBlock(content string) string {
	return fmt.Sprintf("$$\n%s$$\n", content)
}

func (m *MD) WriteTexBlock(content string) *MD {
	m.Write(TexBlock(content))
	return m
}

func Mention(name string) string {
	return fmt.Sprintf(`@%s`, name)
}

func (m *MD) WriteMention(name string) *MD {
	m.Write(Mention(name))
	return m
}

func (m *MD) WriteMentionLn(name string) *MD {
	m.WriteMention(name)
	m.Writeln()
	return m
}

func Cite(messageID string) string {
	return fmt.Sprintf(`%s/messages/%s`, qapi.Endpoint, messageID)
}

func (m *MD) WriteCite(messageID string) *MD {
	m.Writeln(Cite(messageID))
	return m
}

func File(fileID string) string {
	return fmt.Sprintf(`%s/files/%s`, qapi.Endpoint, fileID)
}

func (m *MD) WriteFile(fileID string) *MD {
	m.Writeln(File(fileID))
	return m
}
