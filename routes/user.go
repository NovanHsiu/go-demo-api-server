package routes

import (
	"github.com/NovanHsiu/go-demo-api-server/controllers"
	"github.com/NovanHsiu/go-demo-api-server/middlewares"
	"github.com/NovanHsiu/go-demo-api-server/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUserGroup(db *gorm.DB, config utils.Config, apiGroup *gin.RouterGroup) {
	user := controllers.UserController{DB: db, Config: config}
	userGroup := apiGroup.Group("/users")
	{
		// no restriction
		userGroup.POST("/login", user.LogIn)
		// authorized by token
		userGroup.Use(middlewares.AuthSessionToken(db))
		userGroup.DELETE("/logout", user.LogOut)
		userGroup.GET("/personalProfile", user.GetUserProfile)
		userGroup.PUT("/personalProfile", user.ModifyUserProfile)
		userGroup.PUT("/personalProfile/password", user.ModifyUserProfilePassword)
		// admin only
		userGroup.Use(middlewares.AuthSessionToken(db)).Use(middlewares.AdminOnly(db))
		userGroup.POST("/", user.AddUser)
		userGroup.GET("/", user.GetUserList)
		userGroup.GET("/:id", user.GetUser)
		userGroup.PUT("/:id", user.ModifyUser)
		userGroup.PUT("/:id/password", user.ModifyUserPassword)
		userGroup.DELETE("/:id", user.DeleteUser)
	}
}
