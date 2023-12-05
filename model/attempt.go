package model

import "gorm.io/gorm"

type Attempt struct {
	gorm.Model
	RemainingAttempt int `gorm:"not null"`
}
