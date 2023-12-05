package model

import "gorm.io/gorm"

type Attempt struct {
	gorm.Model
	RemainingAttempt int `binding:"required,min=0" gorm:"not null"`
}
