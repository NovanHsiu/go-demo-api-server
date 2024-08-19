package gorm

import (
	"context"
	"fmt"

	"github.com/NovanHsiu/go-demo-api-server/internal/adapter/repository/gorm/model"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/common"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/parameter"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/response"
)

func (r *GormRepository) GetUserByID(ctx context.Context, id string) (*response.UserResponseListItem, common.Error) {
	user := model.User{}
	if err := r.db.Where("id = ?", id).Preload("UserRole").Last(&user).Error; err != nil {
		return nil, common.NewError(common.ErrorCodeRemoteProcess, err)
	} else if user.ID == 0 {
		return nil, common.NewError(common.ErrorCodeResourceNotFound, fmt.Errorf("user not found"))
	}
	resp := user.GetResponse()
	return &resp, nil
}

func (r *GormRepository) GetUserAndPasswordByAccount(ctx context.Context, account string) (*response.UserResponseListItem, string, common.Error) {
	user := model.User{}
	if err := r.db.Where("account = ?", account).Preload("UserRole").Last(&user).Error; err != nil {
		return nil, "", common.NewError(common.ErrorCodeRemoteProcess, err)
	} else if user.ID == 0 {
		return nil, "", common.NewError(common.ErrorCodeResourceNotFound, fmt.Errorf("user not found"))
	}
	resp := user.GetResponse()
	return &resp, user.Password, nil
}

func (r *GormRepository) GetUsersByParams(ctx context.Context, params parameter.GetUserList) ([]response.UserResponseListItem, int, common.Error) {
	// set filter
	filterMap := map[string]string{
		"account": params.Account,
		"name":    params.Name,
		"email":   params.Email,
	}
	db := r.db
	for key, value := range filterMap {
		if value != "" {
			db = db.Where(key+" like ?", "%"+value+"%")
		}
	}
	if params.UserRoleName != "" {
		db = db.Joins("inner join user_roles on user_roles.id=users.user_role_id and user_roles.name like ?", "%"+params.UserRoleName+"%")
	}
	// set order
	if params.OrderBy == "" {
		db = db.Order("users.id " + params.Order)
	} else {
		db = db.Order("users." + params.OrderBy + " " + params.Order)
	}
	// set pagination
	userList := []model.User{}
	if params.PageNumber == 0 {
		params.PageNumber = 1
	}
	if params.PageSize == 0 {
		params.PageSize = 10
	}
	db.Offset(params.Page.GetOffset()).Limit(params.PageSize).Preload("UserRole").Find(&userList)
	data := []response.UserResponseListItem{}
	for _, user := range userList {
		data = append(data, user.GetResponse())
	}
	// get pages
	var count int64
	db.Model(&model.User{}).Count(&count)
	return data, params.Page.GetPages(int(count)), nil
}

func (r *GormRepository) CreateUserByParams(ctx context.Context, params parameter.AddUser) (*response.UserResponseListItem, common.Error) {
	// check password and password2
	if params.Password != params.Password2 {
		return nil, common.NewError(common.ErrorCodeParameterInvalid, fmt.Errorf("password and password2 do not match"))
	}
	// check account exists or not
	var count int64
	if r.db.Model(&model.User{}).Where("account=?", params.Account).Count(&count); count > 0 {
		return nil, common.NewError(common.ErrorCodeParameterInvalid, fmt.Errorf("this account is existed"))
	}
	// check user role
	userRole := model.UserRole{}
	if r.db.Where("name=?", params.UserRoleName).Last(&userRole); userRole.ID == 0 {
		return nil, common.NewError(common.ErrorCodeParameterInvalid, fmt.Errorf("userRoleName not found"))
	}
	user := model.User{
		Account:    params.Account,
		Password:   common.Cipher.EncodePassword(params.Password),
		Name:       params.Name,
		Email:      params.Email,
		UserRoleID: userRole.ID,
	}
	if err := r.db.Create(&user).Error; err != nil {
		return nil, common.NewError(common.ErrorCodeRemoteProcess, err)
	}
	resp := user.GetResponse()
	return &resp, nil
}

func (r *GormRepository) UpdateUserByIDnParams(ctx context.Context, id string, params parameter.ModifyUser) common.Error {
	// set update map
	updateMap := params.ModifyPersonalProfile.GetUpdateMap()
	userRole := model.UserRole{}
	if params.UserRoleName != "" {
		if r.db.Where("name=?", params.UserRoleName).Last(&userRole); userRole.ID == 0 {
			return common.NewError(common.ErrorCodeParameterInvalid, fmt.Errorf("user_role.name='%s' not found", params.UserRoleName))
		}
		updateMap["user_role_id"] = userRole.ID
	}
	if err := r.db.Model(&model.User{}).Where("id=?", id).Updates(updateMap).Error; err != nil {
		return common.NewError(common.ErrorCodeRemoteProcess, err)
	}
	return nil
}

func (r *GormRepository) UpdateUserPasswordByIDnParams(ctx context.Context, id string, params parameter.ModifyUserPassword) common.Error {
	user := model.User{}
	if r.db.Where("id=?", id).Last(&user); user.ID == 0 {
		return common.NewError(common.ErrorCodeResourceNotFound, fmt.Errorf("user not found"))
	}
	// check old passowrd is correct
	if !common.Cipher.ComparePassword(user.Password, params.OldPassword) {
		return common.NewError(common.ErrorCodeParameterInvalid, fmt.Errorf("incorrected old password"))
	}
	if err := r.db.Model(&model.User{}).Where("id=?", id).Update("password", common.Cipher.EncodePassword(params.NewPassword2)).Error; err != nil {
		return common.NewError(common.ErrorCodeRemoteProcess, err)
	}
	return nil
}

func (r *GormRepository) DeleteUserByID(ctx context.Context, id string) common.Error {
	if err := r.db.Where("id=?", id).Delete(&model.User{}).Error; err != nil {
		return common.NewError(common.ErrorCodeRemoteProcess, err)
	}
	return nil
}
