package md

import "strings"

func Header(text string, level int) string {
	return strings.Repeat("#", level) + " " + text
}

func (m *MD) WriteHeader(text string, level int) *MD {
	m.Writeln(Header(text, level))
	return m
}

func H1(text string) string {
	return Header(text, 1)
}

func (m *MD) WriteH1(text string) *MD {
	m.Writeln(H1(text))
	return m
}

func H2(text string) string {
	return Header(text, 2)
}

func (m *MD) WriteH2(text string) *MD {
	m.Writeln(H2(text))
	return m
}

func H3(text string) string {
	return Header(text, 3)
}

func (m *MD) WriteH3(text string) *MD {
	m.Writeln(H3(text))
	return m
}

func H4(text string) string {
	return Header(text, 4)
}

func (m *MD) WriteH4(text string) *MD {
	m.Writeln(H4(text))
	return m
}

func H5(text string) string {
	return Header(text, 5)
}

func (m *MD) WriteH5(text string) *MD {
	m.Writeln(H5(text))
	return m
}

func H6(text string) string {
	return Header(text, 6)
}

func (m *MD) WriteH6(text string) *MD {
	m.Writeln(H6(text))
	return m
}
