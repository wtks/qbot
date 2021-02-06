package md

import "fmt"

func (m *MD) WriteSimpleList(texts []string) *MD {
	for _, text := range texts {
		m.Writeln("+", text)
	}
	return m
}

func (m *MD) WriteSimpleNumList(texts []string) *MD {
	for i, text := range texts {
		m.Writeln(fmt.Sprintf("%d.", i+1), text)
	}
	return m
}
