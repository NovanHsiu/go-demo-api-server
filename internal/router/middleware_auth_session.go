package router

import (
	"fmt"
	"time"

	"github.com/NovanHsiu/go-demo-api-server/internal/domain/common"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/constant"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (m *Middleware) AuthSessionToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		var token interface{}
		if cacheToken, ok := m.App.Cache.SessionTokenCache.Get(session.ID()); ok {
			token = cacheToken
		} else {
			token = session.Get("token")
			if token != nil {
				m.App.Cache.SessionTokenCache.Add(session.ID(), token, 10*time.Minute)
			}
		}
		if token == nil {
			respondWithError(c, common.NewError(common.ErrorCodeAuthNotAuthenticated, fmt.Errorf("token undefined")))
			c.Abort()
			return
		}
		userID := token.(string)
		c.Set(constant.UserIDKey, userID)
		c.Next()
	}
}
