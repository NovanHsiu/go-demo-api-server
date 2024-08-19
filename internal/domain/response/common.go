package response

type JSONResult struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"ok"`
}

type JSONResultData struct {
	JSONResult
	Data interface{} `json:"data"`
}

type JSONResultDataList struct {
	JSONResultData
	// 總頁數
	Pages int `json:"pages" example:"10"`
}

type IDData struct {
	ID uint `json:"id" example:"1"`
}
