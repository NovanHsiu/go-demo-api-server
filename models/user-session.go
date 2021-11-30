package models

import "gorm.io/gorm"

type UserSession struct {
	gorm.Model
	Token  string
	UserID uint `gorm:"index:user_sessions_users_id_idx"`
	User   User `gorm:"constraint:OnUpdate:SET NULL,OnDelete:SET NULL;"`
}
