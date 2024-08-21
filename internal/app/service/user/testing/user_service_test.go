package testing

import (
	"context"
	"testing"

	"github.com/NovanHsiu/go-demo-api-server/internal/app/service/user"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/common"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/parameter"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/response"
	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUserService_Login(t *testing.T) {
	t.Parallel()
	// Args
	type Args struct {
		Account      string
		Password     string
		UserListItem response.UserResponseListItem
	}
	var args Args
	_ = faker.FakeData(&args)
	args.UserListItem.Account = args.Account

	// Init
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Test cases
	type testCase struct {
		Name         string
		SetupService func(t *testing.T) *user.UserService
		ExpectError  bool
	}

	testCases := []testCase{
		{
			Name:        "login successful",
			ExpectError: false,
			SetupService: func(t *testing.T) *user.UserService {
				mock := buildServiceMock(ctrl)
				encodePassword := common.Cipher.EncodePassword(args.Password)
				mock.UserRepo.EXPECT().GetUserAndPasswordByAccount(gomock.Any(), args.Account).Return(&args.UserListItem, encodePassword, nil)

				service := buildService(mock)
				return service
			},
		},
		{
			Name:        "login failed",
			ExpectError: true,
			SetupService: func(t *testing.T) *user.UserService {
				mock := buildServiceMock(ctrl)

				mock.UserRepo.EXPECT().GetUserAndPasswordByAccount(gomock.Any(), args.Account).Return(nil, "", common.DomainError{})

				service := buildService(mock)
				return service
			},
		},
		{
			Name:        "login password error",
			ExpectError: true,
			SetupService: func(t *testing.T) *user.UserService {
				mock := buildServiceMock(ctrl)

				mock.UserRepo.EXPECT().GetUserAndPasswordByAccount(gomock.Any(), args.Account).Return(&args.UserListItem, "random password!", nil)

				service := buildService(mock)
				return service
			},
		},
	}

	for i := range testCases {
		c := testCases[i]
		t.Run(c.Name, func(t *testing.T) {
			service := c.SetupService(t)
			param := parameter.Login{
				Account:  args.Account,
				Password: args.Password,
			}
			_, err := service.Login(context.Background(), param)
			if c.ExpectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
