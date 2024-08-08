package gopdf

func NewBorderStyle(top, left, right, bottom bool, color *RGB) *BorderStyle {
	return &BorderStyle{
		Top:    top,
		Left:   left,
		Right:  right,
		Bottom: bottom,
		Color:  color,
	}
}

type BorderStyle struct {
	Top    bool
	Left   bool
	Right  bool
	Bottom bool
	Color  *RGB
}

func (s *BorderStyle) SetupBorderColor(p *PDF) {
	if s.Color != nil {
		p.Engine.SetDrawColor(s.Color.R, s.Color.G, s.Color.B)
	}
}

func (s *BorderStyle) BorderToEngineString() string {
	var str string
	if s.Top {
		str += "T"
	}
	if s.Left {
		str += "L"
	}
	if s.Right {
		str += "R"
	}
	if s.Bottom {
		str += "B"
	}
	return str
}
