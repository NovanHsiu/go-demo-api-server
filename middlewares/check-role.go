package middlewares

import (
	"net/http"

	"github.com/NovanHsiu/go-demo-api-server/models"
	"github.com/NovanHsiu/go-demo-api-server/utils"
	"github.com/NovanHsiu/go-demo-api-server/utils/constants"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminOnly(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString(constants.UserIDKey)
		user := models.User{}
		db.Where("id=?", userID).Preload("UserRole").Last(&user)
		if user.UserRole.Code != 1 {
			c.JSON(http.StatusForbidden, utils.GetResponseObject(40301, "no permission"))
			c.Abort()
			return
		}
		c.Next()
	}
}
