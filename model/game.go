package model

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	BoxId    uint   `gorm:"not null" json:"box_id"`
	WalletId uint   `gorm:"not null" json:"wallet_id"`
	Box      Box    `gorm:"foreignKey:box_id;references:id"`
	Wallet   Wallet `gorm:"foreignKey:wallet_id;references:id"`
}
