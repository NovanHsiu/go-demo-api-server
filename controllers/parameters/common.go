package parameters

type Page struct {
	// 排序，預設 `desc`
	Order string `json:"order" default:"desc" enums:"desc,asc"`
	// 頁碼，預設 1
	PageNumber int `json:"pageNumber" default:"1"`
	// 每頁資料筆數，預設 10
	PageSize int `json:"pageSize" default:"10"`
}

func (p *Page) GetOffset() int {
	return (p.PageNumber - 1) * p.PageSize
}

func (p *Page) GetPages(elementCount int) int {
	if elementCount%p.PageSize == 0 {
		return elementCount / p.PageSize
	} else {
		return (elementCount / p.PageSize) + 1
	}
}
