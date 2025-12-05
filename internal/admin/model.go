package admin

import "gorm.io/gorm"

type Admin struct {
	*gorm.Model
	Username string `gorm:"unique;not null" json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
