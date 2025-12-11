package admin

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	    ID        uint           `gorm:"primaryKey" json:"id"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Username string `gorm:"unique;not null" json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role string `json:"role" gorm:"check: role IN ('admin','vendedor')" binding:"required"`
}
