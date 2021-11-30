package utils

type JSONResult struct {
	ErrorCode    int    `json:"error_code" example:"20001"`
	ErrorMessage string `json:"error_message" example:"ok"`
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

// GetResponseObject get object for response
// errorCode:
// 20001: 請求成功
// 20101: 創建成功
// 20401: 無內容 (通常為更新或刪除請求回傳結果)
// 40001: Bad Request，缺少必要參數
// 40002: Bad Request，伺服器找不到參數資料
// 40003: Bad Request，伺服器不允許創建重複的特定資訊
// 40101: 需要授權以回應請求
// 40301: 無權限訪問
// 50001: 未特別分類的伺服器內部錯誤
// 50002: 伺服器資料庫錯誤
func GetResponseObject(errorCode int, errorMessage string) JSONResult {
	return JSONResult{
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
	}
}

func GetResponseObjectData(errorCode int, errorMessage string, data interface{}) JSONResultData {
	result := JSONResultData{
		Data: data,
	}
	result.ErrorCode = errorCode
	result.ErrorMessage = errorMessage
	return result
}

func GetResponseObjectDataList(errorCode int, errorMessage string, data interface{}, pages int) JSONResultDataList {
	result := JSONResultData{
		Data: data,
	}
	result.ErrorCode = errorCode
	result.ErrorMessage = errorMessage
	return JSONResultDataList{
		JSONResultData: result,
		Pages:          pages,
	}
}
