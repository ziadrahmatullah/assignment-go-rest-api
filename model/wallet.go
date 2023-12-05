package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	WalletNumber string          `gorm:"not null"`
	Balance      decimal.Decimal `gorm:"not null"`
	UserId       uint            `gorm:"not null"`
	User         User            `gorm:"foreignKey:user_id;references:id"`
}
