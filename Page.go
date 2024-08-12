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
	PageLayout *PageLayout `json:"page_layout"`

	PageHeight       float64 `json:"page_height"`
	PageWidth        float64 `json:"page_width"`
	PageMarginTop    float64 `json:"page_margin_top"`
	PageMarginBottom float64 `json:"page_margin_bottom"`
	PageMarginLeft   float64 `json:"page_margin_left"`
	PageMarginRight  float64 `json:"page_margin_right"`
	PageBodyHeight   float64 `json:"page_body_height"`
	PageBodyWidth    float64 `json:"page_body_width"`
}
