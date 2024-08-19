package user

import (
	"context"

	"github.com/NovanHsiu/go-demo-api-server/internal/domain/common"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/parameter"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/response"
)

//go:generate mockgen -destination automock/good_repository.go -package=automock . GoodRepository
type UserRepository interface {
	GetUserByID(ctx context.Context, id string) (*response.UserResponseListItem, common.Error)
	GetUserAndPasswordByAccount(ctx context.Context, account string) (*response.UserResponseListItem, string, common.Error)
	GetUsersByParams(ctx context.Context, params parameter.GetUserList) ([]response.UserResponseListItem, int, common.Error)
	CreateUserByParams(ctx context.Context, params parameter.AddUser) (*response.UserResponseListItem, common.Error)
	UpdateUserByIDnParams(ctx context.Context, id string, params parameter.ModifyUser) common.Error
	UpdateUserPasswordByIDnParams(ctx context.Context, id string, params parameter.ModifyUserPassword) common.Error
	DeleteUserByID(ctx context.Context, id string) common.Error
}
