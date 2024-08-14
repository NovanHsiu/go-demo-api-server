package middleware

import (
	"net/http"

	"github.com/NovanHsiu/go-demo-api-server/internal/adapter/gorm/model"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/common"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/constant"
	"github.com/gin-gonic/gin"
)

func (m *Middleware) AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString(constant.UserIDKey)
		user := model.User{}
		m.App.DB.Where("id=?", userID).Preload("UserRole").Last(&user)
		if user.UserRole.Code != 1 {
			c.JSON(http.StatusForbidden, common.GetResponseObject(40301, "no permission"))
			c.Abort()
			return
		}
		c.Next()
	}
}
