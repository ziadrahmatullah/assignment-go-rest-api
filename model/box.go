package model

import "gorm.io/gorm"

type Box struct {
	gorm.Model
	RewardAmount int `gorm:"not null"`
}
