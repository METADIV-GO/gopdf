package gopdf

func NewPage(pdf *PDF) *Page {
	page := &Page{PageLayout: pdf.PageLayout}
	pH, pW := pdf.Engine.GetPageSize()
	l, t, r, b := pdf.Engine.GetMargins()
	page.PageHeight = pH
	page.PageWidth = pW
	page.PageMarginTop = t
	page.PageMarginBottom = b
	page.PageMarginLeft = l
	page.PageMarginRight = r
	page.PageBodyHeight = pH - t - b
	page.PageBodyWidth = pW - l - r
	return page
}

type Page struct {
	PageLayout *PageLayout

	PageHeight       float64
	PageWidth        float64
	PageMarginTop    float64
	PageMarginBottom float64
	PageMarginLeft   float64
	PageMarginRight  float64
	PageBodyHeight   float64
	PageBodyWidth    float64
}
