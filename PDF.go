package gopdf

import (
	"bufio"
	"bytes"
	"fmt"

	"github.com/METADIV-GO/gopdf/ttf_bytes"
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
	Engine           *gofpdf.Fpdf
	PageLayout       *PageLayout
	DefaultFontStyle *FontStyle
	CurrentPageIndex int

	PageHeight       float64
	PageWidth        float64
	PageMarginTop    float64
	PageMarginBottom float64
	PageMarginLeft   float64
	PageMarginRight  float64
	PageBodyHeight   float64 // Page height minus top and bottom margins
	PageBodyWidth    float64 // Page width minus left and right margins
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
	p.Engine.WriteAligned(0, style.LineHeight, text, p.processHAlign(align))
}

func (p *PDF) WriteImage(imgSrc string, width float64, height float64) {
	if width == 0 {
		width = p.PageBodyWidth
	}
	p.Engine.Image(imgSrc, p.PageMarginLeft, p.Engine.GetY(), width, height, true, "", 0, "")
}

func (p *PDF) WriteTable(cells []*Cell) {
	numOfCells := len(cells)
	widthOfCell := p.PageBodyWidth / float64(numOfCells)
	for _, cell := range cells {
		style := cell.Style
		style.FontStyle.Setup(p)
		style.SetupFillColor(p)
		style.BorderStyle.SetupBorderColor(p)
		p.Engine.CellFormat(
			widthOfCell,
			style.FontStyle.LineHeight,
			cell.Text,
			cell.Style.BorderStyle.BorderToEngineString(),
			0,
			cell.Style.ToAlignEngineString(),
			cell.Style.FillColor != nil, 0, "")
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
}

func (p *PDF) initDefaultFontStyle() {
	p.DefaultFontStyle = NewFontStyle("", 0, 0, nil, false, false, false)
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
	if len(layout) == 0 {
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
