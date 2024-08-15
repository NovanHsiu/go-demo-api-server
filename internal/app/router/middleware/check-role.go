package middlewares

import (
	"net/http"

	"github.com/NovanHsiu/go-demo-api-server/internal/adapter/repository/gorm/model"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/common"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/constant"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminOnly(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString(constant.UserIDKey)
		user := model.User{}
		db.Where("id=?", userID).Preload("UserRole").Last(&user)
		if user.UserRole.Code != 1 {
			c.JSON(http.StatusForbidden, common.GetResponseObject(40301, "no permission"))
			c.Abort()
			return
		}
		c.Next()
	}
}
