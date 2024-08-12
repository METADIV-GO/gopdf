package gopdf

func NewCell(text string, style *CellStyle, width float64, widthPercentage float64) *Cell {
	c := &Cell{
		Text:         text,
		Width:        width,
		WidthPercent: widthPercentage,
	}
	c.SetStyle(style)
	return c
}

type Cell struct {
	Text         string     `json:"text"`
	Style        *CellStyle `json:"style"`
	Width        float64    `json:"width"`
	WidthPercent float64    `json:"width_percent"`
}

func (c *Cell) SetStyle(style *CellStyle) {
	if style != nil {
		c.Style = style
	} else {
		c.Style = NewCellStyle(nil, nil, nil, "", "")
	}
}
