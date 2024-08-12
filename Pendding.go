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
	Top    float64 `json:"top"`
	Left   float64 `json:"left"`
	Right  float64 `json:"right"`
	Bottom float64 `json:"bottom"`
}
