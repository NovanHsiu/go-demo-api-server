package router

import (
	"errors"
	"log"

	"github.com/NovanHsiu/go-demo-api-server/internal/domain/common"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/response"
	"github.com/gin-gonic/gin"
)

type ErrorMessage struct {
	Code         int    `json:"code"`
	ErrorMessage string `json:"error_message"`
	Description  string `json:"description"`
}

func respondWithError(c *gin.Context, err error) {
	errMessage := parseError(err)
	ctx := c.Request.Context()
	log.Println(ctx, errMessage)
	c.AbortWithStatusJSON(errMessage.Code, errMessage)
}

func parseError(err error) ErrorMessage {
	var domainError common.DomainError
	// We don't check if errors.As is valid or not
	// because an empty common.DomainError would return default error data.
	_ = errors.As(err, &domainError)

	return ErrorMessage{
		Code:         domainError.StatusCode(),
		ErrorMessage: domainError.ErrorMessage,
		Description:  domainError.Description,
	}
}

func respondJsonResult(c *gin.Context, code int, message string) {
	c.JSON(code, response.JSONResult{
		Code:    code,
		Message: message,
	})
}

func respondWithData(c *gin.Context, code int, message string, data interface{}) {
	result := response.JSONResultData{
		Data: data,
	}
	result.Code = code
	result.Message = message
	c.JSON(code, result)
}

func respondWithDataList(c *gin.Context, code int, message string, data interface{}, pages int) {
	result := response.JSONResultData{
		Data: data,
	}
	result.Code = code
	result.Message = message
	c.JSON(code, response.JSONResultDataList{
		JSONResultData: result,
		Pages:          pages,
	})
}
