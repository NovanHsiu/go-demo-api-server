package route

import (
	"os"

	_ "github.com/NovanHsiu/go-demo-api-server/docs"
	"github.com/NovanHsiu/go-demo-api-server/routes"
	"github.com/NovanHsiu/go-demo-api-server/utils"
	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gorm.io/gorm"
)

func GetRoutingEngine(db *gorm.DB, config utils.Config) *gin.Engine {
	eng := gin.Default()
	eng.Use(gin.Recovery())
	// allow CORS
	cconfig := cors.DefaultConfig()
	cconfig.AllowAllOrigins = true
	cconfig.AllowCredentials = true
	cconfig.AllowedHeaders = append(cconfig.AllowedHeaders, []string{"Authorization"}...)
	eng.Use(cors.New(cconfig))
	// "yoursecretpassowrd" is password for encoding
	store := gormsessions.NewStore(db, true, []byte("yoursecretpassowrd"))
	// set session middleware "mysession" is session and cookie name
	// store is storage engine, we can use redis or another db to store session
	eng.Use(sessions.Sessions("api-server-session", store))

	// set swagger url
	eng.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // api-docs/index.html

	// set middleware

	// set route
	os.Mkdir(utils.GetExecutionDir()+"/"+config.File.StaticFileDir, os.ModePerm)
	eng.Use(static.Serve("/static", static.LocalFile(config.File.StaticFileDir, true)))
	apiGroup := eng.Group("/api")
	// /users
	routes.SetUserGroup(db, config, apiGroup)
	return eng
}
