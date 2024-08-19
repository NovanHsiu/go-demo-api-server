package router

import "github.com/NovanHsiu/go-demo-api-server/internal/app"

type Middleware struct {
	App *app.Application
}
