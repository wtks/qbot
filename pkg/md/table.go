package md

import "strings"

type Table struct {
	rowSize int
	colSize int
	body    [][]string
}

func (t *Table) SetHeader(col int, content string) *Table {
	if col < t.colSize {
		t.body[0][col] = content
	}
	return t
}

func (t *Table) SetHeaders(contents ...string) *Table {
	for i, content := range contents {
		t.SetHeader(i, content)
	}
	return t
}

func (t *Table) SetContent(row, col int, content string) *Table {
	if row < t.rowSize && col < t.colSize {
		row = row + 2
		t.body[row][col] = content
	}
	return t
}

func (t *Table) SetCols(row int, contents ...string) *Table {
	for col, content := range contents {
		t.SetContent(row, col, content)
	}
	return t
}

func (t *Table) String() string {
	var buffer strings.Builder
	for _, row := range t.body {
		buffer.WriteString("|")
		for _, col := range row {
			buffer.WriteString(col)
			buffer.WriteString("|")
		}
		buffer.WriteString("\n")

	}
	return buffer.String()
}

func NewTable(row, col int) *Table {
	t := &Table{
		colSize: col,
		rowSize: row,
	}
	row = row + 2
	t.body = make([][]string, row)
	for i := 0; i < row; i++ {
		t.body[i] = make([]string, col)
		if i == 1 {
			for j := 0; j < col; j++ {
				t.body[i][j] = "----"
			}
		}
	}
	return t
}

func (m *MD) WriteTable(t *Table) *MD {
	m.Write(t.String())
	return m
}
