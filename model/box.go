package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Box struct {
	gorm.Model
	RewardAmount decimal.Decimal `gorm:"not null"`
}
