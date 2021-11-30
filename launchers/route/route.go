package route

import (
	"os"

	"github.com/NovanHsiu/go-demo-api-server/routes"
	"github.com/NovanHsiu/go-demo-api-server/utils"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gorm.io/gorm"
)

func GetRoutingEngine(db *gorm.DB, config utils.Config) *gin.Engine {
	eng := gin.New()
	eng.Use(gin.Recovery())
	// allow CORS
	cconfig := cors.DefaultConfig()
	cconfig.AllowAllOrigins = true
	cconfig.AllowCredentials = true
	cconfig.AllowedHeaders = append(cconfig.AllowedHeaders, []string{"Authorization"}...)
	eng.Use(cors.New(cconfig))

	// set swagger url
	eng.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // api-docs/index.html

	// set middleware

	// set route
	os.Mkdir(utils.ExecutionDir()+"/"+config.File["static_file_dir"], os.ModePerm)
	eng.Use(static.Serve("/static", static.LocalFile(config.File["static_file_dir"], true)))
	apiGroup := eng.Group("/api")
	routes.SetUserGroup(db, config, apiGroup)
	return eng
}
