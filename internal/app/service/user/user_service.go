package user

import (
	"context"
	"fmt"

	"github.com/NovanHsiu/go-demo-api-server/internal/domain/common"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/parameter"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/response"
)

func (s *UserService) Login(ctx context.Context, params parameter.Login) (*response.UserResponseListItem, common.Error) {
	item, password, err := s.userRepo.GetUserAndPasswordByAccount(ctx, params.Account)
	if err != nil {
		return nil, err
	}
	if !common.Cipher.ComparePassword(password, params.Password) {
		return nil, common.NewError(common.ErrorCodeParameterInvalid, fmt.Errorf("password error"))
	}
	return item, err
}

func (s *UserService) GetUser(ctx context.Context, userID string) (*response.UserResponseListItem, common.Error) {
	return s.userRepo.GetUserByID(ctx, userID)
}

func (s *UserService) GetUserList(ctx context.Context, params parameter.GetUserList) ([]response.UserResponseListItem, int, common.Error) {
	return s.userRepo.GetUsersByParams(ctx, params)
}

func (s *UserService) CreateUser(ctx context.Context, params parameter.AddUser) (*response.UserResponseListItem, common.Error) {
	return s.userRepo.CreateUserByParams(ctx, params)
}

func (s *UserService) UpdateUser(ctx context.Context, userID string, params parameter.ModifyUser) common.Error {
	return s.userRepo.UpdateUserByIDnParams(ctx, userID, params)
}

func (s *UserService) UpdateUserPasswordByIDnParams(ctx context.Context, userID string, params parameter.ModifyUserPassword) common.Error {
	return s.userRepo.UpdateUserPasswordByIDnParams(ctx, userID, params)
}

func (s *UserService) DeleteUser(ctx context.Context, userID string) common.Error {
	return s.userRepo.DeleteUserByID(ctx, userID)
}
