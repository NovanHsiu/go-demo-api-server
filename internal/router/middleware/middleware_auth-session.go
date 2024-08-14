package middleware

import (
	"net/http"

	"github.com/NovanHsiu/go-demo-api-server/internal/domain/common"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/constant"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (m *Middleware) AuthSessionToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		token := session.Get("token")
		if token == nil {
			c.JSON(http.StatusUnauthorized, common.GetResponseObject(40101, "token undefined"))
			c.Abort()
			return
		}
		userID := token.(string)
		c.Set(constant.UserIDKey, userID)
		c.Next()
	}
}
