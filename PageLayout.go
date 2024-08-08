package gopdf

func NewPageLayout(orientation, paper string) *PageLayout {
	layout := &PageLayout{}
	layout.SetOrientation(orientation)
	layout.SetPaper(paper)
	return layout
}

type PageLayout struct {
	Orientation      string
	Paper            string
	PageMargin       *PageMargin
	DefaultFontStyle *FontStyle
}

/*
SetOrientation sets the orientation for the PDF.
By default, the orientation is portrait.
*/
func (s *PageLayout) SetOrientation(orientation string) {
	switch orientation {
	case OrientationPortrait, OrientationLandscape:
		s.Orientation = orientation
	default:
		s.Orientation = OrientationPortrait
	}
}

/*
SetPaper sets the paper size for the PDF.
By default, the paper size is A4.
*/
func (s *PageLayout) SetPaper(paper string) {
	switch paper {
	case PaperA3, PaperA4, PaperA5, PaperLetter, PaperLegal, PaperTabloid:
		s.Paper = paper
	default:
		s.Paper = PaperA4
	}
}
