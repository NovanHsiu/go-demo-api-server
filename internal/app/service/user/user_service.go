package user

import (
	"context"
	"fmt"

	"github.com/NovanHsiu/go-demo-api-server/internal/domain/common"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/parameter"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/response"
)

func (s *UserService) Login(ctx context.Context, param parameter.Login) (*response.UserResponseListItem, common.Error) {
	item, password, err := s.userRepo.GetUserAndPasswordByAccount(ctx, param.Account)
	if err != nil {
		return nil, err
	}
	if !common.Cipher.ComparePassword(password, param.Password) {
		return nil, common.NewError(common.ErrorCodeParameterInvalid, fmt.Errorf("password error"))
	}
	return item, err
}

func (s *UserService) GetUser(ctx context.Context, userID string) (*response.UserResponseListItem, common.Error) {
	return s.userRepo.GetUserByID(ctx, userID)
}

func (s *UserService) GetUserList(ctx context.Context, param parameter.GetUserList) ([]response.UserResponseListItem, int, common.Error) {
	return s.userRepo.GetUsersByParam(ctx, param)
}

func (s *UserService) CreateUser(ctx context.Context, param parameter.AddUser) (*response.UserResponseListItem, common.Error) {
	return s.userRepo.CreateUserByParam(ctx, param)
}

func (s *UserService) UpdateUser(ctx context.Context, userID string, param parameter.ModifyUser) common.Error {
	return s.userRepo.UpdateUserByIDnParam(ctx, userID, param)
}

func (s *UserService) UpdateUserPasswordByIDnParam(ctx context.Context, userID string, param parameter.ModifyUserPassword) common.Error {
	return s.userRepo.UpdateUserPasswordByIDnParam(ctx, userID, param)
}

func (s *UserService) DeleteUser(ctx context.Context, userID string) common.Error {
	return s.userRepo.DeleteUserByID(ctx, userID)
}
