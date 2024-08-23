package testing

import (
	"context"
	"testing"

	"github.com/NovanHsiu/go-demo-api-server/internal/domain/parameter"
	"github.com/stretchr/testify/assert"
)

func TestGormRepository_GetUserByID(t *testing.T) {
	db := testDB
	repo := initRepository(db)
	user, err := repo.GetUserByID(context.Background(), "1")
	assert.NoError(t, err)
	assert.Equal(t, uint(1), user.ID)
}

func TestGormRepository_GetUserAndPasswordByAccount(t *testing.T) {
	db := testDB
	repo := initRepository(db)
	user, passwd, err := repo.GetUserAndPasswordByAccount(context.Background(), "admin")
	assert.NoError(t, err)
	assert.Equal(t, "admin", user.Account)
	assert.Greater(t, len(passwd), 1)
}

func TestGormRepository_GetUsersByParam(t *testing.T) {
	db := testDB
	repo := initRepository(db)
	users, pages, err := repo.GetUsersByParam(context.Background(), parameter.GetUserList{
		Account: "admin",
	})
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, pages, 1)
	assert.Equal(t, "admin", users[0].Account)
	users, _, err = repo.GetUsersByParam(context.Background(), parameter.GetUserList{
		Name: "管理者",
	})
	assert.NoError(t, err)
	assert.Equal(t, "admin", users[0].Account)
	users, _, err = repo.GetUsersByParam(context.Background(), parameter.GetUserList{
		UserRoleName: "管理者",
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, users[0].UserRole.Code)
	users, _, err = repo.GetUsersByParam(context.Background(), parameter.GetUserList{
		Email: "admin@",
	})
	assert.NoError(t, err)
	assert.Equal(t, "admin", users[0].Account)
	users, _, err = repo.GetUsersByParam(context.Background(), parameter.GetUserList{})
	assert.NoError(t, err)
	assert.Len(t, users, 2)
}

func TestGormRepository_CreateUserByParam(t *testing.T) {
	db := testDB
	repo := initRepository(db)
	testAccount := "createdtest"
	user, err := repo.CreateUserByParam(context.Background(), parameter.AddUser{
		Account:      testAccount,
		Name:         testAccount,
		Password:     testAccount,
		Password2:    testAccount,
		UserRoleName: "一般使用者",
		Email:        "createdtest@testmail.com",
	})
	assert.NoError(t, err)
	assert.Equal(t, testAccount, user.Account)
	assert.Equal(t, testAccount, user.Name)
}

func TestGormRepository_UpdateUserByIDnParam(t *testing.T) {
	db := testDB
	repo := initRepository(db)
	param := parameter.ModifyUser{
		ModifyPersonalProfile: parameter.ModifyPersonalProfile{
			Name:  "updatedtest",
			Email: "updatedtest@testmail.com",
		},
		UserRoleName: "管理者",
	}
	err := repo.UpdateUserByIDnParam(context.Background(), "2", param)
	assert.NoError(t, err)
	user, err := repo.GetUserByID(context.Background(), "2")
	assert.NoError(t, err)
	assert.Equal(t, param.Name, user.Name)
	assert.Equal(t, param.Email, user.Email)
	assert.Equal(t, param.UserRoleName, user.UserRole.Name)
}

func TestGormRepository_UpdateUserPasswordByIDnParam(t *testing.T) {
	db := testDB
	repo := initRepository(db)
	err := repo.UpdateUserPasswordByIDnParam(context.Background(), "2", parameter.ModifyUserPassword{
		OldPassword: "ttt",
	})
	assert.Error(t, err)
	err = repo.UpdateUserPasswordByIDnParam(context.Background(), "123", parameter.ModifyUserPassword{
		OldPassword: "testuser",
		NewPassword: "testuser",
	})
	assert.Error(t, err)
	err = repo.UpdateUserPasswordByIDnParam(context.Background(), "2", parameter.ModifyUserPassword{
		OldPassword: "testuser",
		NewPassword: "testuser",
	})
	assert.NoError(t, err)
}

func TestGormRepository_DeleteUserByID(t *testing.T) {
	db := testDB
	repo := initRepository(db)
	err := repo.DeleteUserByID(context.Background(), "2")
	assert.NoError(t, err)
	_, err = repo.GetUserByID(context.Background(), "2")
	assert.Error(t, err)
	err = db.Table("users").Where("id = 2").UpdateColumn("deleted_at", nil).Error
	assert.NoError(t, err)
}
