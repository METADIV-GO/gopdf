package gopdf

func NewRGB(r, g, b int) RGB {
	return RGB{
		R: r,
		G: g,
		B: b,
	}
}

type RGB struct {
	R int
	G int
	B int
}
