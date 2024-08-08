package gopdf

func NewPageMargin(top, left, right, bottom float64) *PageMargin {
	return &PageMargin{
		Top:    top,
		Left:   left,
		Right:  right,
		Bottom: bottom,
	}
}

type PageMargin struct {
	Top    float64
	Left   float64
	Right  float64
	Bottom float64
}
