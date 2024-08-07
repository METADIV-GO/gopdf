package gopdf

func NewFontStyle(
	fontFamily string,
	fontSize float64,
	lineHeight float64,
	fontColor *RGB,
	bold, strikeout, underline bool,
) *FontStyle {
	fs := &FontStyle{}
	fs.SetFontFamily(fontFamily)
	fs.SetFontSize(fontSize)
	fs.SetLineHeight(lineHeight)
	fs.SetFontColor(fontColor)
	fs.SetBold(bold)
	fs.SetStrikeout(strikeout)
	fs.SetUnderline(underline)
	return fs
}

type FontStyle struct {
	FontFamily string
	FontSize   float64
	LineHeight float64
	FontColor  *RGB
	Bold       bool
	Strikeout  bool
	Underline  bool
}

/*
Setup sets the font style for the PDF.
*/
func (s *FontStyle) Setup(pdf *PDF) {
	styleStr := ""
	if s.Bold {
		styleStr += "B"
	}
	if s.Strikeout {
		styleStr += "S"
	}
	if s.Underline {
		styleStr += "U"
	}
	pdf.Engine.SetFont(s.FontFamily, styleStr, s.FontSize)
}

/*
SetFontFamily sets the font family for the PDF.
By default, the font family is Helvetica.
*/
func (s *FontStyle) SetFontFamily(fontFamily string) {
	switch fontFamily {
	case FontFamilyCourier, FontFamilyHelvetica, FontFamilyTimes,
		FontFamilySymbol, FontFamilyNotoSansTC, FontFamilyNotoSansSC:
		s.FontFamily = fontFamily
	default:
		s.FontFamily = FontFamilyHelvetica
	}
}

/*
SetFontSize sets the font size for the PDF.
By default, the font size is 12.
*/
func (s *FontStyle) SetFontSize(fontSize float64) {
	if fontSize <= 0 {
		s.FontSize = 12
	} else {
		s.FontSize = fontSize
	}
}

/*
SetLineHeight sets the line height for the PDF.
By default, the line height same as the font size.
*/
func (s *FontStyle) SetLineHeight(lineHeight float64) {
	if lineHeight <= 0 {
		s.LineHeight = s.FontSize
	} else {
		s.LineHeight = lineHeight
	}
}

/*
SetFontColor sets the font color for the PDF.
By default, the font color is black.
*/
func (s *FontStyle) SetFontColor(fontColor *RGB) {
	if fontColor == nil {
		s.FontColor = &RGB{0, 0, 0}
	} else {
		s.FontColor = fontColor
	}
}

/*
SetBold sets the font style to bold.
By default, the font style is not bold.
*/
func (s *FontStyle) SetBold(bold bool) {
	s.Bold = bold
}

/*
SetStrikeout sets the font style to strikeout.
By default, the font style is not strikeout.
*/
func (s *FontStyle) SetStrikeout(strikeout bool) {
	s.Strikeout = strikeout
}

/*
SetUnderline sets the font style to underline.
By default, the font style is not underline.
*/
func (s *FontStyle) SetUnderline(underline bool) {
	s.Underline = underline
}
