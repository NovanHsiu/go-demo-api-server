package app

import (
	"context"
	"time"

	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"

	adapterGorm "github.com/NovanHsiu/go-demo-api-server/internal/adapter/repository/gorm"
	"github.com/NovanHsiu/go-demo-api-server/internal/app/service/user"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/common"
	"github.com/patrickmn/go-cache"
)

type Application struct {
	ApplicationParam ApplicationParam
	Cache            *ApplicationCache
	SessionsStore    sessions.Store
	UserService      *user.UserService
}

type ApplicationParam struct {
	Config common.Config
}

type ApplicationCache struct {
	SessionTokenCache *cache.Cache
}

func NewApplication(ctx context.Context, param ApplicationParam) (*Application, error) {
	// set db
	db, err := adapterGorm.NewDB(param.Config.DB)
	if err != nil {
		return nil, err
	}
	adapterGorm.CreateDefaultTable(db)
	tokenCache := cache.New(10*time.Minute, 20*time.Minute)
	// "yoursecretpassowrd" is password for encoding
	sessionsStore := gormsessions.NewStore(db, true, []byte("yoursecretpassowrd"))
	repo := adapterGorm.NewGormRepository(ctx, db)
	userService := user.NewUserService(user.UserServiceParam{
		UserRepo: repo,
	})
	app := Application{
		ApplicationParam: param,
		Cache: &ApplicationCache{
			SessionTokenCache: tokenCache,
		},
		SessionsStore: sessionsStore,
		UserService:   userService,
	}
	return &app, nil
}
