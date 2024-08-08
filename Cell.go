package gopdf

func NewCell(text string, style *CellStyle) *Cell {
	c := &Cell{
		Text: text,
	}
	c.SetStyle(style)
	return c
}

type Cell struct {
	Text  string
	Style *CellStyle
}

func (c *Cell) SetStyle(style *CellStyle) {
	if style != nil {
		c.Style = style
	} else {
		c.Style = NewCellStyle(nil, nil, nil, "", "")
	}
}
