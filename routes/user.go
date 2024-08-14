package routes

import (
	"github.com/NovanHsiu/go-demo-api-server/controllers"
	"github.com/NovanHsiu/go-demo-api-server/internal/app"
	"github.com/NovanHsiu/go-demo-api-server/middlewares"
	"github.com/gin-gonic/gin"
)

func SetUserGroup(app *app.Application, apiGroup *gin.RouterGroup) {
	user := controllers.UserController{App: app}
	userGroup := apiGroup.Group("/users")
	{
		// no restriction
		userGroup.POST("/login", user.LogIn)
		// authorized by token
		userGroup.Use(middlewares.AuthSessionToken(app.DB))
		userGroup.DELETE("/logout", user.LogOut)
		userGroup.GET("/personalProfile", user.GetUserProfile)
		userGroup.PUT("/personalProfile", user.ModifyUserProfile)
		userGroup.PUT("/personalProfile/password", user.ModifyUserProfilePassword)
		// admin only
		userGroup.Use(middlewares.AuthSessionToken(app.DB)).Use(middlewares.AdminOnly(app.DB))
		userGroup.POST("/", user.AddUser)
		userGroup.GET("/", user.GetUserList)
		userGroup.GET("/:id", user.GetUser)
		userGroup.PUT("/:id", user.ModifyUser)
		userGroup.PUT("/:id/password", user.ModifyUserPassword)
		userGroup.DELETE("/:id", user.DeleteUser)
	}
}
