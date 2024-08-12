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
	Top    float64 `json:"top"`
	Left   float64 `json:"left"`
	Right  float64 `json:"right"`
	Bottom float64 `json:"bottom"`
}
