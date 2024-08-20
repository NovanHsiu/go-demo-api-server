package user

import (
	"context"

	"github.com/NovanHsiu/go-demo-api-server/internal/domain/common"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/parameter"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/response"
)

//go:generate mockgen -destination automock/user_repository.go -package=automock . UserRepository
type UserRepository interface {
	GetUserByID(ctx context.Context, id string) (*response.UserResponseListItem, common.Error)
	GetUserAndPasswordByAccount(ctx context.Context, account string) (*response.UserResponseListItem, string, common.Error)
	GetUsersByParam(ctx context.Context, param parameter.GetUserList) ([]response.UserResponseListItem, int, common.Error)
	CreateUserByParam(ctx context.Context, param parameter.AddUser) (*response.UserResponseListItem, common.Error)
	UpdateUserByIDnParam(ctx context.Context, id string, param parameter.ModifyUser) common.Error
	UpdateUserPasswordByIDnParam(ctx context.Context, id string, param parameter.ModifyUserPassword) common.Error
	DeleteUserByID(ctx context.Context, id string) common.Error
}
