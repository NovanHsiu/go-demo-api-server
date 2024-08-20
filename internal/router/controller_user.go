package router

import (
	"fmt"
	"net/http"

	"github.com/NovanHsiu/go-demo-api-server/internal/app"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/common"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/constant"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/parameter"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/response"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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
// @Success 200 {object} response.JSONResultData{data=response.UserResponseListItem} "ok"
// @Router /users/login [post]
// LogIn login user's account
func (uc *UserController) LogIn(c *gin.Context) {
	param := parameter.Login{}
	if err := c.ShouldBind(&param); err != nil {
		respondWithError(c, common.NewError(common.ErrorCodeParameterInvalid, err))
		return
	}
	user, err := uc.App.UserService.Login(c.Request.Context(), param)
	if err != nil {
		respondWithError(c, err)
		return
	}
	session := sessions.Default(c)
	uc.App.Cache.SessionTokenCache.Delete(session.ID())
	session.Set("token", fmt.Sprintf("%d", user.ID))
	session.Options(sessions.Options{
		MaxAge: 60 * 60 * 24, // expired time 24 hours
	})
	session.Save()
	respondWithData(c, http.StatusOK, "ok", user)
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
	uc.App.Cache.SessionTokenCache.Delete(session.ID())
	session.Delete("token")
	session.Save()
	respondJsonResult(c, 204, "no content")
}

// @Summary 取得使用者個人資料
// @Tags users
// @Accept  json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.JSONResultData{data=response.UserResponseListItem} "ok"
// @Router /users/personalProfile [get]
// GetProfile get user's profile
func (uc *UserController) GetUserProfile(c *gin.Context) {
	userID := c.GetString(constant.UserIDKey)
	user, err := uc.App.UserService.GetUser(c.Request.Context(), userID)
	if err != nil {
		respondWithError(c, err)
		return
	}
	respondWithData(c, http.StatusOK, "ok", user)
}

// @Summary 新增使用者
// @Tags users
// @Accept  json
// @Produce json
// @Security ApiKeyAuth
// @Param addUserData body parameter.AddUser true "新增使用者資料"
// @Success 201 {object} response.JSONResultData{data=response.IDData} "created"
// @Router  /users [post]
// AddUser add user
func (uc *UserController) AddUser(c *gin.Context) {
	param := parameter.AddUser{}
	if err := c.ShouldBind(&param); err != nil {
		respondWithError(c, common.NewError(common.ErrorCodeParameterInvalid, err))
		return
	}
	user, err := uc.App.UserService.CreateUser(c.Request.Context(), param)
	if err != nil {
		respondWithError(c, err)
		return
	}
	respondWithData(c, 201, "created", response.IDData{ID: user.ID})
}

// @Summary 取得使用者資料清單
// @Description 取得使用者資料清單，不輸入過濾參數則回傳所有用戶資料。
// @Tags users
// @Accept  json
// @Produce json
// @Security ApiKeyAuth
// @Param getUserList query parameter.GetUserList true "取得使用者過濾資訊"
// @Success 200 {object} response.JSONResultDataList{data=[]response.UserResponseListItem} "successful operation"
// @Router /users [get]
// GetUserList get a list of user's data
func (uc *UserController) GetUserList(c *gin.Context) {
	param := parameter.GetUserList{}
	if err := c.ShouldBind(&param); err != nil {
		respondWithError(c, common.NewError(common.ErrorCodeParameterInvalid, err))
		return
	}
	userList, pages, err := uc.App.UserService.GetUserList(c.Request.Context(), param)
	if err != nil {
		respondWithError(c, err)
		return
	}
	respondWithDataList(c, 200, "ok", userList, pages)
}

// @Summary 取得使用者資料
// @Tags users
// @Accept  json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.JSONResultData{data=response.UserResponseListItem} "ok"
// @Param id path int true "使用者ID"
// @Router /users/{id} [get]
// GetUser get specificed user's profile
func (uc *UserController) GetUser(c *gin.Context) {
	user, err := uc.App.UserService.GetUser(c.Request.Context(), c.Param("id"))
	if err != nil {
		respondWithError(c, err)
		return
	}
	respondWithData(c, http.StatusOK, "ok", user)
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
	param := parameter.ModifyUser{}
	if err := c.ShouldBind(&param); err != nil {
		respondWithError(c, common.NewError(common.ErrorCodeParameterInvalid, err))
		return
	}
	if err := uc.App.UserService.UpdateUser(c.Request.Context(), c.Param("id"), param); err != nil {
		respondWithError(c, err)
		return
	}
	respondJsonResult(c, 200, "ok")
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
	param := parameter.ModifyUserPassword{}
	if err := c.ShouldBind(&param); err != nil {
		respondWithError(c, common.NewError(common.ErrorCodeParameterInvalid, err))
		return
	}
	if err := uc.App.UserService.UpdateUserPasswordByIDnParam(c.Request.Context(), c.Param("id"), param); err != nil {
		respondWithError(c, err)
		return
	}
	respondJsonResult(c, 204, "no content")
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
	if err := uc.App.UserService.DeleteUser(c.Request.Context(), c.Param("id")); err != nil {
		respondWithError(c, err)
		return
	}
	respondJsonResult(c, 204, "no content")
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
	param := parameter.ModifyPersonalProfile{}
	if err := c.ShouldBind(&param); err != nil {
		respondWithError(c, common.NewError(common.ErrorCodeParameterInvalid, err))
		return
	}
	userID := c.GetString(constant.UserIDKey)
	if err := uc.App.UserService.UpdateUser(c.Request.Context(), userID, parameter.ModifyUser{
		ModifyPersonalProfile: param,
	}); err != nil {
		respondWithError(c, err)
		return
	}
	respondJsonResult(c, 204, "no content")
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
	param := parameter.ModifyUserPassword{}
	if err := c.ShouldBind(&param); err != nil {
		respondWithError(c, common.NewError(common.ErrorCodeParameterInvalid, err))
		return
	}
	userID := c.GetString(constant.UserIDKey)
	if err := uc.App.UserService.UpdateUserPasswordByIDnParam(c.Request.Context(), userID, param); err != nil {
		respondWithError(c, err)
		return
	}
	respondJsonResult(c, 204, "no content")
}
