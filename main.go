package main

import (
	"context"
	"fmt"
	"os"

	"github.com/NovanHsiu/go-demo-api-server/internal/app"
	routeLauncher "github.com/NovanHsiu/go-demo-api-server/launchers/route"
	"github.com/NovanHsiu/go-demo-api-server/utils"
)

// @title Go Demo API Server
// @version 1.1.0
// @description ## 摘要
// @description 可用來做為 GO API Server 教學展示或 API 服務基礎模板
// @description ## Swagger API 認證
// @description 以 `[POST] /users/login` API 取得token，點選文件頁面右側 Authorize 按鈕輸入token 作為Value儲存認證

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
			fmt.Println("version 1.1.0")
			return
		}
	}
	//log.SetOutput(os.Stdout)
	config := utils.GetConfig()
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
	routeEng := routeLauncher.GetRoutingEngine(app)
	// run routing enigine
	if config.Common.TlsCrtPath != "" && config.Common.TlsKeyPath != "" {
		// https
		if _, err := os.Stat(config.Common.TlsCrtPath); !os.IsNotExist(err) {
			go routeEng.RunTLS(":"+sslport, config.Common.TlsCrtPath, config.Common.TlsKeyPath)
		}
	}
	routeEng.Run(":" + port)
}
