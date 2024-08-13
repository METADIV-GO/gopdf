package gopdf

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"

	"github.com/METADIV-GO/gopdf/ttf_bytes"
	"github.com/METADIV-GO/nanoid"
	"github.com/jung-kurt/gofpdf"
)

func New(layout ...*PageLayout) *PDF {
	pdf := new(PDF)
	pdf.CurrentPageIndex = -1
	pdf.PageLayout = pdf.processLayoutOpts(layout...)
	pdf.initEngine(pdf.PageLayout)
	pdf.initDefaultSupportingFonts()
	pdf.initDefaultFontStyle()
	pdf.initPageBodySize()
	pdf.Engine.AddPage()
	return pdf
}

type PDF struct {
	Engine           *gofpdf.Fpdf `json:"-"`
	PageLayout       *PageLayout  `json:"page_layout"`
	DefaultFontStyle *FontStyle   `json:"default_font_style"`
	CurrentPageIndex int          `json:"-"`

	PageHeight       float64 `json:"page_height"`
	PageWidth        float64 `json:"page_width"`
	PageMarginTop    float64 `json:"page_margin_top"`
	PageMarginBottom float64 `json:"page_margin_bottom"`
	PageMarginLeft   float64 `json:"page_margin_left"`
	PageMarginRight  float64 `json:"page_margin_right"`
	PageBodyHeight   float64 `json:"page_body_height"` // Page height minus top and bottom margins
	PageBodyWidth    float64 `json:"page_body_width"`  // Page width minus left and right margins
}

func (p *PDF) AddPage() {
	p.Engine.AddPage()
	p.initPageBodySize()
}

func (p *PDF) SetDefaultFontStyle(style *FontStyle) {
	p.DefaultFontStyle = style
}

func (p *PDF) WriteText(text string, style *FontStyle) {
	if style == nil {
		style = p.DefaultFontStyle
	}
	style.Setup(p)
	p.Engine.Write(style.LineHeight, text)
}

func (p *PDF) WriteLink(text string, link string, style *FontStyle) {
	if style == nil {
		style = p.DefaultFontStyle
	}
	p.Engine.WriteLinkString(style.LineHeight, text, link)
}

func (p *PDF) WriteTextBox(text string, align string, style *FontStyle) {
	if style == nil {
		style = p.DefaultFontStyle
	}
	style.Setup(p)
	lines := strings.Split(text, "\n")
	for i := range lines {
		p.Engine.WriteAligned(0, style.LineHeight, lines[i], p.processHAlign(align))
		p.LineBreak(style)
	}
}

func (p *PDF) WriteImage(imgSrc string, width float64, height float64, flow bool) {
	if width == 0 {
		width = p.PageBodyWidth
	}
	p.Engine.Image(imgSrc, p.PageMarginLeft, p.Engine.GetY(), width, height, flow, "", 0, "")
}

func (p *PDF) WriteImageBytes(imgBytes []byte, width float64, height float64, flow bool) {
	if width == 0 {
		width = p.PageBodyWidth
	}
	name := nanoid.NewSafe()
	p.Engine.RegisterImageOptionsReader(name, gofpdf.ImageOptions{ImageType: "png"}, bytes.NewReader(imgBytes))
	p.Engine.Image(name, p.PageMarginLeft, p.Engine.GetY(), width, height, flow, "", 0, "")
}

func (p *PDF) WriteTable(cells []*Cell, padding *Padding) {

	if padding != nil {
		if padding.Top > 0 {
			p.Engine.Ln(padding.Top)
		}
		if padding.Left > 0 {
			p.Engine.SetX(p.PageMarginLeft + padding.Left)
		}
	}

	pageBodyWidth := p.PageBodyWidth
	if padding != nil && padding.Right > 0 {
		pageBodyWidth -= padding.Right
	}
	if padding != nil && padding.Left > 0 {
		pageBodyWidth -= padding.Left
	}

	var emptyWidth int
	var usedWidth float64
	for i := range cells {
		if cells[i].Width == 0 && cells[i].WidthPercent == 0 {
			emptyWidth++
		} else if cells[i].WidthPercent > 0 && cells[i].Width == 0 {
			cells[i].Width = pageBodyWidth * cells[i].WidthPercent
			usedWidth += cells[i].Width
		} else {
			usedWidth += cells[i].Width
		}
	}

	widthOfCell := (pageBodyWidth - usedWidth) / float64(emptyWidth)
	for i := range cells {
		if cells[i].Width == 0 {
			cells[i].Width = widthOfCell
		}
	}

	for _, cell := range cells {
		style := cell.Style
		style.FontStyle.Setup(p)
		style.SetupFillColor(p)
		style.BorderStyle.SetupBorderColor(p)
		p.Engine.MultiCell(
			cell.Width,
			style.FontStyle.LineHeight,
			cell.Text,
			cell.Style.BorderStyle.BorderToEngineString(),
			cell.Style.ToAlignEngineString(),
			cell.Style.FillColor != nil)
	}
	if padding != nil && padding.Bottom > 0 {
		p.Engine.Ln(padding.Bottom)
	}
}

func (p *PDF) LineBreak(style *FontStyle) {
	if style == nil {
		style = p.DefaultFontStyle
	}
	p.Engine.Ln(style.LineHeight)
}

func (p *PDF) ToFile(filePath string) {
	err := p.Engine.OutputFileAndClose(filePath)
	if err != nil {
		fmt.Println(err)
	}
}

func (p *PDF) ToBytes() []byte {
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	p.Engine.Output(writer)
	err := writer.Flush()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return b.Bytes()
}

func (p *PDF) initEngine(layout *PageLayout) {
	p.Engine = gofpdf.New(layout.Orientation, "pt", layout.Paper, "")
	if layout.PageMargin != nil {
		p.Engine.SetMargins(layout.PageMargin.Left, layout.PageMargin.Top, layout.PageMargin.Right)
	}
}

func (p *PDF) initDefaultFontStyle() {
	if p.PageLayout.DefaultFontStyle == nil {
		p.DefaultFontStyle = NewFontStyle("", 0, 0, nil, false, false, false)
	} else {
		p.DefaultFontStyle = p.PageLayout.DefaultFontStyle
	}
}

func (p *PDF) initPageBodySize() {
	pW, pH := p.Engine.GetPageSize()
	l, t, r, b := p.Engine.GetMargins()
	p.PageHeight = pH
	p.PageWidth = pW
	p.PageMarginTop = t
	p.PageMarginBottom = b
	p.PageMarginLeft = l
	p.PageMarginRight = r
	p.PageBodyHeight = pH - t - b
	p.PageBodyWidth = pW - l - r
}

func (p *PDF) processLayoutOpts(layout ...*PageLayout) *PageLayout {
	if len(layout) == 0 || layout[0] == nil {
		return NewPageLayout(OrientationPortrait, PaperA4)
	}
	return layout[0]
}

func (p *PDF) processHAlign(align string) string {
	switch align {
	case AlignLeft:
		return "L"
	case AlignCenter:
		return "C"
	case AlignRight:
		return "R"
	default:
		return "L"
	}
}

func (p *PDF) initDefaultSupportingFonts() {
	p.Engine.AddUTF8FontFromBytes(FontFamilyNotoSansTC, "", ttf_bytes.NotoSansTCRegular)
	p.Engine.AddUTF8FontFromBytes(FontFamilyNotoSansTC, "B", ttf_bytes.NotoSansTCBold)
	p.Engine.AddUTF8FontFromBytes(FontFamilyNotoSansSC, "", ttf_bytes.NotoSansSCRegular)
	p.Engine.AddUTF8FontFromBytes(FontFamilyNotoSansSC, "B", ttf_bytes.NotoSansSCBold)
}
