package model

import (
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/response"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Account    string
	Password   string
	Name       string
	Email      string
	UserRoleID uint     `gorm:"index:users_user_roles_id_idx"`
	UserRole   UserRole `gorm:"constraint:OnUpdate:SET NULL,OnDelete:SET NULL;"`
}

func (u *User) GetResponse() response.UserResponseListItem {
	return response.UserResponseListItem{
		ID:       u.ID,
		Account:  u.Account,
		Name:     u.Name,
		Email:    u.Email,
		UserRole: response.UserRole(u.UserRole),
	}
}
