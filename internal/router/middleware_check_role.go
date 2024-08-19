package router

import (
	"fmt"

	"github.com/NovanHsiu/go-demo-api-server/internal/domain/common"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/constant"
	"github.com/gin-gonic/gin"
)

func (m *Middleware) AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString(constant.UserIDKey)
		user, err := m.App.UserService.GetUser(c.Request.Context(), userID)
		if err != nil {
			respondWithError(c, err)
			c.Abort()
			return
		}
		if user.UserRole.Code != 1 {
			respondWithError(c, common.NewError(common.ErrorCodeAuthPermissionDenied, fmt.Errorf("no permission")))
			c.Abort()
			return
		}
		c.Next()
	}
}
