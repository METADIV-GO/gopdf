package gopdf

func NewPadding(top, left, right, bottom float64) *Padding {
	return &Padding{
		Top:    top,
		Left:   left,
		Right:  right,
		Bottom: bottom,
	}
}

type Padding struct {
	Top    float64 `json:"top"`
	Left   float64 `json:"left"`
	Right  float64 `json:"right"`
	Bottom float64 `json:"bottom"`
}
