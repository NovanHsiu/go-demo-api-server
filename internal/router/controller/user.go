package controller

import (
	"fmt"
	"net/http"

	"github.com/NovanHsiu/go-demo-api-server/internal/adapter/gorm/model"
	"github.com/NovanHsiu/go-demo-api-server/internal/app"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/common"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/constant"
	"github.com/NovanHsiu/go-demo-api-server/internal/router/controller/parameter"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	App *app.Application
}

// @Summary 登入使用者帳號
// @Description 登入使用者帳號，回傳 Header Set-Cookie 帶有登入 Token 資訊
// @Tags users
// @Accept  json
// @Produce json
// @Param loginData body parameter.Login true "登入資料"
// @Success 200 {object} common.JSONResultData{data=model.UserResponseListData} "ok"
// @Router /users/login [post]
// LogIn login user's account
func (uc *UserController) LogIn(c *gin.Context) {
	params := parameter.Login{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, common.GetResponseObject(40001, err.Error()))
		return
	}
	user := model.User{}
	if err := uc.App.DB.Where("account=?", params.Account).Last(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, common.GetResponseObject(50002, err.Error()))
		return
	}
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, common.GetResponseObject(40002, fmt.Sprintf("can not found account=%s", user.Account)))
		return
	} else if !common.Cipher.ComparePassword(user.Password, params.Password) {
		c.JSON(http.StatusBadRequest, common.GetResponseObject(40002, "password error"))
		return
	}
	session := sessions.Default(c)
	session.Set("token", fmt.Sprintf("%d", user.ID))
	session.Options(sessions.Options{
		MaxAge: 60 * 60 * 24, // expired time 24 hours
	})
	session.Save()
	data := user.GetResponse()
	result := common.GetResponseObjectData(20001, "ok", data)
	c.JSON(http.StatusOK, result)
}

// @Summary 登出使用者帳號
// @Tags users
// @Accept  json
// @Produce json
// @Security ApiKeyAuth
// @Success 204 {string} string "no content"
// @Router /users/logout [delete]
// LogOut logout user's account
func (uc *UserController) LogOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("token")
	session.Save()
	c.JSON(http.StatusNoContent, common.GetResponseObject(20401, "no content"))
}

// @Summary 取得使用者個人資料
// @Tags users
// @Accept  json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} common.JSONResultData{data=model.UserResponseListData} "ok"
// @Router /users/personalProfile [get]
// GetProfile get user's profile
func (uc *UserController) GetUserProfile(c *gin.Context) {
	userID := c.GetString(constant.UserIDKey)
	user := model.User{}
	uc.App.DB.Where("id=?", userID).Preload("UserRole").Last(&user)
	c.JSON(http.StatusOK, common.GetResponseObjectData(20001, "ok", user.GetResponse()))
}

// @Summary 新增使用者
// @Tags users
// @Accept  json
// @Produce json
// @Security ApiKeyAuth
// @Param addUserData body parameter.AddUser true "新增使用者資料"
// @Success 201 {object} common.JSONResultData "created"
// @Router  /users [post]
// AddUser add user
func (uc *UserController) AddUser(c *gin.Context) {
	params := parameter.AddUser{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, common.GetResponseObject(40001, err.Error()))
		return
	}
	// check password and password2
	if params.Password != params.Password2 {
		c.JSON(http.StatusBadRequest, common.GetResponseObject(40001, "password and password2 do not match"))
		return
	}
	// check account exists or not
	var count int64
	if uc.App.DB.Model(&model.User{}).Where("account=?", params.Account).Count(&count); count > 0 {
		c.JSON(http.StatusBadRequest, common.GetResponseObject(40003, "this account is existed"))
		return
	}
	// check user role
	userRole := model.UserRole{}
	if uc.App.DB.Where("name=?", params.UserRoleName).Last(&userRole); userRole.ID == 0 {
		c.JSON(http.StatusBadRequest, common.GetResponseObject(40002, "userRoleName not found"))
		return
	}
	if err := uc.App.DB.Create(&model.User{
		Account:    params.Account,
		Password:   common.Cipher.EncodePassword(params.Password),
		Name:       params.Name,
		Email:      params.Email,
		UserRoleID: userRole.ID,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, common.GetResponseObject(50002, err.Error()))
		return
	}
	c.JSON(http.StatusCreated, common.GetResponseObject(20101, "created"))
}

// @Summary 取得使用者資料清單
// @Description 取得使用者資料清單，不輸入過濾參數則回傳所有用戶資料。
// @Tags users
// @Accept  json
// @Produce json
// @Security ApiKeyAuth
// @Param getUserList query parameter.GetUserList true "取得使用者過濾資訊"
// @Success 200 {object} common.JSONResultDataList{data=[]model.UserResponseListData} "successful operation"
// @Router /users [get]
// GetUserList get a list of user's data
func (uc *UserController) GetUserList(c *gin.Context) {
	params := parameter.GetUserList{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, common.GetResponseObject(40001, err.Error()))
		return
	}
	// set filter
	filterMap := map[string]string{
		"account": params.Account,
		"name":    params.Name,
		"email":   params.Email,
	}
	db := uc.App.DB
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
	db.Offset(params.Page.GetOffset()).Limit(params.PageSize).Preload("UserRole").Find(&userList)
	data := []model.UserResponseListData{}
	for _, user := range userList {
		data = append(data, user.GetUserResponseListData())
	}
	// get pages
	var count int64
	db.Model(&model.User{}).Count(&count)
	c.JSON(http.StatusOK, common.GetResponseObjectDataList(20001, "ok", data, params.Page.GetPages(int(count))))
}

// @Summary 取得使用者資料
// @Tags users
// @Accept  json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} common.JSONResultData{data=model.UserResponseListData} "ok"
// @Param id path int true "使用者ID"
// @Router /users/{id} [get]
// GetUser get specificed user's profile
func (uc *UserController) GetUser(c *gin.Context) {
	user := model.User{}
	if uc.App.DB.Where("id=?", c.Param("id")).Preload("UserRole").Last(&user); user.ID == 0 {
		c.JSON(http.StatusBadRequest, common.GetResponseObject(40002, "user not found"))
		return
	}
	c.JSON(http.StatusOK, common.GetResponseObjectData(20001, "ok", user.GetUserResponseListData()))
}

// @Summary 修改使用者資料
// @Tags users
// @Accept  json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "使用者ID"
// @Param modifyUserData body parameter.ModifyUser true "修改使用者資料"
// @Success 204 {string} string "no content"
// @Router /users/{id} [put]
// ModifyUser modify specificed user's profile
func (uc *UserController) ModifyUser(c *gin.Context) {
	params := parameter.ModifyUser{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, common.GetResponseObject(40001, err.Error()))
		return
	}
	// set update map
	updateMap := params.ModifyPersonalProfile.GetUpdateMap()
	userRole := model.UserRole{}
	if params.UserRoleName != "" {
		if uc.App.DB.Where("name=?", params.UserRoleName).Last(&userRole); userRole.ID == 0 {
			c.JSON(http.StatusBadRequest, common.GetResponseObject(40002, fmt.Sprintf("user_role.name='%s' not found", params.UserRoleName)))
			return
		}
		updateMap["user_role_id"] = userRole.ID
	}
	uc.App.DB.Model(&model.User{}).Where("id=?", c.Param("id")).Updates(updateMap)
	c.JSON(http.StatusNoContent, common.GetResponseObject(20401, "no content"))
}

// @Summary 修改使用者密碼
// @Tags users
// @Accept  json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "使用者ID"
// @Param modifyUserPassword body parameter.ModifyUserPassword true "修改使用者密碼資料"
// @Success 204 {string} string "no content"
// @Router /users/{id}/password [put]
// ModifyUserPassword modify specificed user's password
func (uc *UserController) ModifyUserPassword(c *gin.Context) {
	params := parameter.ModifyUserPassword{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if statusCode, errResult := modifyPassword(uc.App.DB, params, c.Param("id")); statusCode != http.StatusOK {
		c.JSON(statusCode, errResult)
		return
	}
	c.JSON(http.StatusNoContent, common.GetResponseObject(20401, "no content"))
}

// @Summary 刪除使用者
// @Tags users
// @Accept  json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "使用者ID"
// @Success 204 {string} string "no content"
// @Router /users/{id} [delete]
// DeleteUser delete specificed user
func (uc *UserController) DeleteUser(c *gin.Context) {
	if err := uc.App.DB.Where("id=?", c.Param("id")).Delete(&model.User{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, common.GetResponseObject(50002, err.Error()))
		return
	}
	c.JSON(http.StatusNoContent, common.GetResponseObject(20401, "no content"))
}

// @Summary 修改使用者個人資料
// @Tags users
// @Accept  json
// @Produce json
// @Security ApiKeyAuth
// @Param modifyPersonalProfile body parameter.ModifyPersonalProfile true "修改使用者個人資料"
// @Success 204 {string} string "no content"
// @Router /users/personalProfile [put]
// ModifyUserProfile modify user's profile
func (uc *UserController) ModifyUserProfile(c *gin.Context) {
	params := parameter.ModifyPersonalProfile{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(400, err.Error())
		return
	}
	userID := c.GetString(constant.UserIDKey)
	// set update map
	updateMap := params.GetUpdateMap()
	if err := uc.App.DB.Model(&model.User{}).Where("id=?", userID).Updates(updateMap).Error; err != nil {
		c.JSON(http.StatusInternalServerError, common.GetResponseObject(50002, err.Error()))
		return
	}
	c.JSON(http.StatusNoContent, common.GetResponseObject(20401, "no content"))
}

// @Summary 修改使用者個人密碼
// @Tags users
// @Accept  json
// @Produce json
// @Security ApiKeyAuth
// @Param modifyUserPassword body parameter.ModifyUserPassword true "修改使用者個人密碼資料"
// @Success 204 {string} string "no content data"
// @Router /users/personalProfile/password [put]
// ModifyUserPassword modify specificed user's password
func (uc *UserController) ModifyUserProfilePassword(c *gin.Context) {
	params := parameter.ModifyUserPassword{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(400, err.Error())
		return
	}
	userID := c.GetString(constant.UserIDKey)
	if statusCode, errResult := modifyPassword(uc.App.DB, params, userID); statusCode != http.StatusOK {
		c.JSON(statusCode, errResult)
		return
	}
	c.JSON(http.StatusNoContent, common.GetResponseObject(20401, "no content"))
}

func modifyPassword(db *gorm.DB, params parameter.ModifyUserPassword, userID string) (int, common.JSONResult) {
	user := model.User{}
	if db.Where("id=?", userID).Last(&user); user.ID == 0 {
		return http.StatusBadRequest, common.GetResponseObject(40002, "user not found")
	}
	// check old passowrd is correct
	if !common.Cipher.ComparePassword(user.Password, params.OldPassword) {
		return http.StatusBadRequest, common.GetResponseObject(40002, "incorrected old password")
	}
	// check password and password2
	if params.NewPassword != params.NewPassword2 {
		return http.StatusBadRequest, common.GetResponseObject(40001, "new password and password2 do not match")
	}
	if err := db.Model(&model.User{}).Where("id=?", userID).Update("password", common.Cipher.EncodePassword(params.NewPassword2)).Error; err != nil {
		return http.StatusBadRequest, common.GetResponseObject(50002, err.Error())
	}
	return http.StatusOK, common.JSONResult{}
}
