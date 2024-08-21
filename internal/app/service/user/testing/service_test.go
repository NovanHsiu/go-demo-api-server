package testing

import (
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"

	"github.com/NovanHsiu/go-demo-api-server/internal/app/service/user"
	"github.com/NovanHsiu/go-demo-api-server/internal/app/service/user/automock"
)

type serviceMock struct {
	UserRepo *automock.MockUserRepository
}

func buildServiceMock(ctrl *gomock.Controller) serviceMock {
	return serviceMock{
		UserRepo: automock.NewMockUserRepository(ctrl),
	}
}
func buildService(mock serviceMock) *user.UserService {
	param := user.UserServiceParam{
		UserRepo: mock.UserRepo,
	}
	return user.NewUserService(param)
}

// nolint
func TestMain(m *testing.M) {
	// To avoid getting an empty object slice
	_ = faker.SetRandomMapAndSliceMinSize(2)

	// To avoid getting a zero random number
	_ = faker.SetRandomNumberBoundaries(1, 100)

	m.Run()
}
