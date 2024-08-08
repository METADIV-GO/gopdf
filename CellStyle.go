package gopdf

func NewCellStyle(fontStyle *FontStyle, borderStyle *BorderStyle, fillColor *RGB, hAlign, vAlign string) *CellStyle {
	style := &CellStyle{}
	style.SetFontStyle(fontStyle)
	style.SetBorderStyle(borderStyle)
	style.SetFillColor(fillColor)
	style.SetHAlign(hAlign)
	style.SetVAlign(vAlign)
	return style
}

type CellStyle struct {
	FontStyle   *FontStyle
	BorderStyle *BorderStyle
	FillColor   *RGB
	HAlign      string
	VAlign      string
}

func (s *CellStyle) SetFontStyle(style *FontStyle) {
	if style != nil {
		s.FontStyle = style
	} else {
		s.FontStyle = NewFontStyle("", 0, 0, nil, false, false, false)
	}
}

func (s *CellStyle) SetBorderStyle(style *BorderStyle) {
	if style != nil {
		s.BorderStyle = style
	} else {
		s.BorderStyle = NewBorderStyle(false, false, false, false, nil)
	}
}

func (s *CellStyle) SetFillColor(color *RGB) {
	s.FillColor = color
}

func (s *CellStyle) SetHAlign(align string) {
	switch align {
	case AlignLeft, AlignCenter, AlignRight:
		s.HAlign = align
	default:
		s.HAlign = AlignLeft
	}
}

func (s *CellStyle) SetVAlign(align string) {
	switch align {
	case AlignTop, AlignMiddle, AlignBottom, AlignBaseline:
		s.VAlign = align
	default:
		s.VAlign = AlignTop
	}
}

func (s *CellStyle) SetupFillColor(p *PDF) {
	if s.FillColor != nil {
		p.Engine.SetFillColor(s.FillColor.R, s.FillColor.G, s.FillColor.B)
	}
}

func (s *CellStyle) ToAlignEngineString() string {
	return s.HAlignToEngineString() + s.VAlignToEngineString()
}

func (s *CellStyle) HAlignToEngineString() string {
	switch s.HAlign {
	case AlignLeft:
		return "L"
	case AlignCenter:
		return "C"
	case AlignRight:
		return "R"
	default:
		return "L"
	}
}

func (s *CellStyle) VAlignToEngineString() string {
	switch s.VAlign {
	case AlignTop:
		return "T"
	case AlignMiddle:
		return "M"
	case AlignBottom:
		return "B"
	case AlignBaseline:
		return "A"
	default:
		return "T"
	}
}
