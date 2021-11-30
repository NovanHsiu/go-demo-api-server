package middlewares

import (
	"net/http"

	"github.com/NovanHsiu/go-demo-api-server/models"
	"github.com/NovanHsiu/go-demo-api-server/utils"
	"github.com/NovanHsiu/go-demo-api-server/utils/constants"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthToken(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		var count int64
		if db.Model(&models.UserSession{}).Where("token=?", token).Count(&count); count == 0 {
			c.JSON(http.StatusUnauthorized, utils.GetResponseObject(40101, "token not found"))
			c.Abort()
			return
		}
		userID, err := utils.Cipher.ParseJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.GetResponseObject(40101, err.Error()))
			c.Abort()
			return
		}
		c.Set(constants.UserIDKey, userID)
		c.Set("token", token)
		c.Next()
	}
}
