package model

import (
	"time"

	"gorm.io/gorm"
)

type ResetPassToken struct {
	gorm.Model
	Token  string    `gorm:"not null"`
	Expire time.Time `gorm:"not null"`
	IsUsed bool      `gorm:"not null"`
	UserId uint      `gorm:"not null" json:"user_id"`
	User   User      `gorm:"foreignKey:user_id;references:id"`
}
