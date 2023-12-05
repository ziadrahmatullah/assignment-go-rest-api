package model

import "gorm.io/gorm"

type Attempt struct {
	gorm.Model
	WalletId         uint   `gorm:"not null" json:"wallet_id"`
	RemainingAttempt int    `binding:"required,min=0" gorm:"not null"`
	Wallet           Wallet `gorm:"foreignKey:wallet_id;references:id"`
}
