package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Account    string
	Password   string
	Name       string
	Email      string
	UserRoleID uint     `gorm:"index:users_user_roles_id_idx"`
	UserRole   UserRole `gorm:"constraint:OnUpdate:SET NULL,OnDelete:SET NULL;"`
}

type UserResponseListData struct {
	ID       uint     `json:"id" example:"1"`
	Account  string   `json:"account" example:"admin"`
	Name     string   `json:"name" example:"管理者"`
	Email    string   `json:"email" example:"admin@testmail.com"`
	UserRole UserRole `json:"userRole"`
}

func (u *User) GetResponse() UserResponseListData {
	return UserResponseListData{
		ID:       u.ID,
		Account:  u.Account,
		Name:     u.Name,
		Email:    u.Email,
		UserRole: u.UserRole,
	}
}

func (u *User) GetUserResponseListData() UserResponseListData {
	return UserResponseListData{
		ID:       u.ID,
		Account:  u.Account,
		Name:     u.Name,
		Email:    u.Email,
		UserRole: u.UserRole,
	}
}
