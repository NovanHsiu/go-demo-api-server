package response

type UserRole struct {
	ID         uint   `json:"id" example:"1"`
	Code       int    `gorm:"uniqueIndex" json:"code" example:"1"`
	Name       string `gorm:"uniqueIndex" json:"name" example:"管理者"`
	Permission string `json:"permission" example:"{}"`
}

type UserResponseListItem struct {
	ID       uint     `json:"id" example:"1"`
	Account  string   `json:"account" example:"admin"`
	Name     string   `json:"name" example:"管理者"`
	Email    string   `json:"email" example:"admin@testmail.com"`
	UserRole UserRole `json:"userRole"`
}
