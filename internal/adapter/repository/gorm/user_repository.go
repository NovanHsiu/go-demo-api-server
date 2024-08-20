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

func (r *GormRepository) GetUsersByParam(ctx context.Context, param parameter.GetUserList) ([]response.UserResponseListItem, int, common.Error) {
	// set filter
	filterMap := map[string]string{
		"account": param.Account,
		"name":    param.Name,
		"email":   param.Email,
	}
	db := r.db
	for key, value := range filterMap {
		if value != "" {
			db = db.Where(key+" like ?", "%"+value+"%")
		}
	}
	if param.UserRoleName != "" {
		db = db.Joins("inner join user_roles on user_roles.id=users.user_role_id and user_roles.name like ?", "%"+param.UserRoleName+"%")
	}
	// set order
	if param.OrderBy == "" {
		db = db.Order("users.id " + param.Order)
	} else {
		db = db.Order("users." + param.OrderBy + " " + param.Order)
	}
	// set pagination
	userList := []model.User{}
	if param.PageNumber == 0 {
		param.PageNumber = 1
	}
	if param.PageSize == 0 {
		param.PageSize = 10
	}
	db.Offset(param.Page.GetOffset()).Limit(param.PageSize).Preload("UserRole").Find(&userList)
	data := []response.UserResponseListItem{}
	for _, user := range userList {
		data = append(data, user.GetResponse())
	}
	// get pages
	var count int64
	db.Model(&model.User{}).Count(&count)
	return data, param.Page.GetPages(int(count)), nil
}

func (r *GormRepository) CreateUserByParam(ctx context.Context, param parameter.AddUser) (*response.UserResponseListItem, common.Error) {
	// check password and password2
	if param.Password != param.Password2 {
		return nil, common.NewError(common.ErrorCodeParameterInvalid, fmt.Errorf("password and password2 do not match"))
	}
	// check account exists or not
	var count int64
	if r.db.Model(&model.User{}).Where("account=?", param.Account).Count(&count); count > 0 {
		return nil, common.NewError(common.ErrorCodeParameterInvalid, fmt.Errorf("this account is existed"))
	}
	// check user role
	userRole := model.UserRole{}
	if r.db.Where("name=?", param.UserRoleName).Last(&userRole); userRole.ID == 0 {
		return nil, common.NewError(common.ErrorCodeParameterInvalid, fmt.Errorf("userRoleName not found"))
	}
	user := model.User{
		Account:    param.Account,
		Password:   common.Cipher.EncodePassword(param.Password),
		Name:       param.Name,
		Email:      param.Email,
		UserRoleID: userRole.ID,
	}
	if err := r.db.Create(&user).Error; err != nil {
		return nil, common.NewError(common.ErrorCodeRemoteProcess, err)
	}
	resp := user.GetResponse()
	return &resp, nil
}

func (r *GormRepository) UpdateUserByIDnParam(ctx context.Context, id string, param parameter.ModifyUser) common.Error {
	// set update map
	updateMap := param.ModifyPersonalProfile.GetUpdateMap()
	userRole := model.UserRole{}
	if param.UserRoleName != "" {
		if r.db.Where("name=?", param.UserRoleName).Last(&userRole); userRole.ID == 0 {
			return common.NewError(common.ErrorCodeParameterInvalid, fmt.Errorf("user_role.name='%s' not found", param.UserRoleName))
		}
		updateMap["user_role_id"] = userRole.ID
	}
	if err := r.db.Model(&model.User{}).Where("id=?", id).Updates(updateMap).Error; err != nil {
		return common.NewError(common.ErrorCodeRemoteProcess, err)
	}
	return nil
}

func (r *GormRepository) UpdateUserPasswordByIDnParam(ctx context.Context, id string, param parameter.ModifyUserPassword) common.Error {
	user := model.User{}
	if r.db.Where("id=?", id).Last(&user); user.ID == 0 {
		return common.NewError(common.ErrorCodeResourceNotFound, fmt.Errorf("user not found"))
	}
	// check old passowrd is correct
	if !common.Cipher.ComparePassword(user.Password, param.OldPassword) {
		return common.NewError(common.ErrorCodeParameterInvalid, fmt.Errorf("incorrected old password"))
	}
	if err := r.db.Model(&model.User{}).Where("id=?", id).Update("password", common.Cipher.EncodePassword(param.NewPassword2)).Error; err != nil {
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
