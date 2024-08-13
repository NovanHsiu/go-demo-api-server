package middlewares

import (
	"net/http"

	"github.com/NovanHsiu/go-demo-api-server/utils"
	"github.com/NovanHsiu/go-demo-api-server/utils/constants"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthSessionToken(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		token := session.Get("token")
		if token == nil {
			c.JSON(http.StatusUnauthorized, utils.GetResponseObject(40101, "token undefined"))
			c.Abort()
			return
		}
		userID := token.(string)
		c.Set(constants.UserIDKey, userID)
		c.Next()
	}
}
