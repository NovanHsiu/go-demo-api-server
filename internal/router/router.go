package router

import (
	"os"

	_ "github.com/NovanHsiu/go-demo-api-server/docs"
	"github.com/NovanHsiu/go-demo-api-server/internal/app"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/common"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func GetRouterEngine(app *app.Application) *gin.Engine {
	eng := gin.Default()
	eng.Use(gin.Recovery())
	// allow CORS
	cconfig := cors.DefaultConfig()
	cconfig.AllowAllOrigins = true
	cconfig.AllowCredentials = true
	cconfig.AllowedHeaders = append(cconfig.AllowedHeaders, []string{"Authorization"}...)
	eng.Use(cors.New(cconfig))
	// set session middleware "mysession" is session and cookie name
	// store is storage engine, we can use redis or another db to store session
	eng.Use(sessions.Sessions("api-server-session", app.SessionsStore))

	// set swagger url
	eng.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // api-docs/index.html

	// init middleware
	middleware := Middleware{App: app}
	// set route
	config := app.ApplicationParam.Config
	os.Mkdir(common.GetExecutionDir()+"/"+config.File.StaticFileDir, os.ModePerm)
	eng.Use(static.Serve("/static", static.LocalFile(config.File.StaticFileDir, true)))
	apiGroup := eng.Group("/api")
	// /users
	user := UserController{App: app}
	userGroup := apiGroup.Group("/users")
	{
		// no restriction
		userGroup.POST("/login", user.LogIn)
		// authorized by token
		userGroup.Use(middleware.AuthSessionToken())
		userGroup.DELETE("/logout", user.LogOut)
		userGroup.GET("/personalProfile", user.GetUserProfile)
		userGroup.PUT("/personalProfile", user.ModifyUserProfile)
		userGroup.PUT("/personalProfile/password", user.ModifyUserProfilePassword)
		// admin only
		userGroup.Use(middleware.AuthSessionToken()).Use(middleware.AdminOnly())
		userGroup.POST("/", user.AddUser)
		userGroup.GET("/", user.GetUserList)
		userGroup.GET("/:id", user.GetUser)
		userGroup.PUT("/:id", user.ModifyUser)
		userGroup.PUT("/:id/password", user.ModifyUserPassword)
		userGroup.DELETE("/:id", user.DeleteUser)
	}
	return eng
}
