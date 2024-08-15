package main

import (
	"context"
	"fmt"
	"os"

	"github.com/NovanHsiu/go-demo-api-server/internal/app"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/common"
	"github.com/NovanHsiu/go-demo-api-server/internal/router"
)

// @title Go Demo API Server
// @version 1.2.4
// @description ## 摘要
// @description 可用來做為 GO API Server 教學展示或 API 服務基礎模板
// @description ## Swagger API 認證
// @description 以 `[POST] /users/login` 登入會進行 Session 認證，瀏覽器會自行存取 Cookie 並使用 Session Token。

// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email huey_yu@acebiotek.com
// @license.name MIT License
// @license.url https://opensource.org/licenses/MIT

// @host localhost:3000
// @BasePath /api
// @schemes http https

// @tag.name users
// @tag.description 使用者
func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "version" {
			fmt.Println("version 1.2.2")
			return
		}
	}
	//log.SetOutput(os.Stdout)
	config := common.GetConfig()
	app, err := app.NewApplication(context.Background(), app.ApplicationParams{
		Config: config,
	})
	if err != nil {
		panic(err)
	}
	// set route
	port := os.Getenv("PORT")
	sslport := config.Common.SslPort
	if len(port) == 0 {
		port = config.Common.Port
	}
	routeEng := router.GetRouterEngine(app)
	// run routing enigine
	if config.Common.TlsCrtPath != "" && config.Common.TlsKeyPath != "" {
		// https
		if _, err := os.Stat(config.Common.TlsCrtPath); !os.IsNotExist(err) {
			go routeEng.RunTLS(":"+sslport, config.Common.TlsCrtPath, config.Common.TlsKeyPath)
		}
	}
	routeEng.Run(":" + port)
}
