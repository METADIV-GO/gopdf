package gopdf

func NewRGB(r, g, b int) *RGB {
	return &RGB{
		R: r,
		G: g,
		B: b,
	}
}

type RGB struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}
