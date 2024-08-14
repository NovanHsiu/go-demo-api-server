package model

type UserRole struct {
	ID         uint   `json:"id" example:"1"`
	Code       int    `gorm:"uniqueIndex" json:"code" example:"1"`
	Name       string `gorm:"uniqueIndex" json:"name" example:"管理者"`
	Permission string `json:"permission" example:"{}"`
}

var DefaultUserRole = []UserRole{
	{ID: 1, Code: 1, Name: "管理者", Permission: "{}"},
	{ID: 2, Code: 999, Name: "一般使用者", Permission: "{}"},
}
