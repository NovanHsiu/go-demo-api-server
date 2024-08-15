package parameter

type Login struct {
	Account  string `binding:"required" json:"account" example:"admin"`
	Password string `binding:"required" json:"password" example:"admin" validate:"max=32,min=4"`
}

type ModifyPersonalProfile struct {
	Name  string `json:"name" example:"王大明"`
	Email string `json:"email" example:"daming_wang@testmail.com"`
}

type ModifyUser struct {
	ModifyPersonalProfile
	UserRoleName string `json:"userRoleName" enums:"管理者,一般使用者"`
}

func (m *ModifyPersonalProfile) GetUpdateMap() map[string]interface{} {
	updateMap := map[string]interface{}{}
	if m.Name != "" {
		updateMap["name"] = m.Name
	}
	if m.Email != "" {
		updateMap["email"] = m.Email
	}
	return updateMap
}

type ModifyUserPassword struct {
	OldPassword  string `binding:"required" validate:"max=32,min=4"`
	NewPassword  string `binding:"required" validate:"max=32,min=4"`
	NewPassword2 string `binding:"required" validate:"max=32,min=4"`
}

type GetUserList struct {
	Page
	// 使用者帳號
	Account string `form:"account"`
	// 使用者名稱
	Name string `form:"name"`
	// 使用者角色名稱
	UserRoleName string `form:"userRoleName"`
	// 使用者電子郵件
	Email string `form:"email"`
	// 排序欄位
	OrderBy string `form:"orderBy" enums:"account,name,email"`
}

type AddUser struct {
	Account      string `binding:"required" json:"account" example:"normal_user"`
	Name         string `binding:"required" json:"name" example:"普通的使用者"`
	UserRoleName string `binding:"required" json:"userRoleName" enums:"管理者,一般使用者"`
	Email        string `binding:"required" json:"email" example:"normal_user@testmail.com" validate:"email"`
	Password     string `binding:"required" json:"password" example:"normalguy123" validate:"max=32,min=4"`
	Password2    string `binding:"required" json:"password2" example:"normalguy123" validate:"max=32,min=4"`
}
