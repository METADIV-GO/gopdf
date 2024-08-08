package gopdf

func NewPending(top, left, right, bottom float64) *Pending {
	return &Pending{
		Top:    top,
		Left:   left,
		Right:  right,
		Bottom: bottom,
	}
}

type Pending struct {
	Top    float64
	Left   float64
	Right  float64
	Bottom float64
}
