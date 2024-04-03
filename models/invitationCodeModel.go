package models

import "gorm.io/gorm"

type InvitationCode struct {
	gorm.Model
	InvitationCode string `gorm:"unique"`
	IsUsed         bool   `gorm:"default:false"`
	UserID         uint
}
