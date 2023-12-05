package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string    `gorm:"not null"`
	Birthdate time.Time `gorm:"not null"`
	Email     string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
}
