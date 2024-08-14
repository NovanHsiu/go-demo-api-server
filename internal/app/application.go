package app

import (
	"context"
	"time"

	adapterGorm "github.com/NovanHsiu/go-demo-api-server/internal/adapter/gorm"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/common"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
)

type Application struct {
	ApplicationParams ApplicationParams
	DB                *gorm.DB
	Cache             *ApplicationCache
}

type ApplicationParams struct {
	Config common.Config
}

type ApplicationCache struct {
	SessionTokenCache *cache.Cache
}

func NewApplication(ctx context.Context, params ApplicationParams) (*Application, error) {
	// set db
	db, err := adapterGorm.NewDB(params.Config.DB)
	if err != nil {
		return nil, err
	}
	adapterGorm.CreateDefaultTable(db)
	tokenCache := cache.New(10*time.Minute, 20*time.Minute)
	app := Application{
		ApplicationParams: params,
		DB:                db,
		Cache: &ApplicationCache{
			SessionTokenCache: tokenCache,
		},
	}
	return &app, nil
}
