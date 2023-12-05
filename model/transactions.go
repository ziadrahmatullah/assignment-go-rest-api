package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type SourceOfFunds string
type TransactionTypes string

const (
	BankTransfer SourceOfFunds = "Bank Transfer"
	CreditCard   SourceOfFunds = "Credit Card"
	Cash         SourceOfFunds = "Cash"
	Reward       SourceOfFunds = "Reward"

	Transfer   TransactionTypes = "Transfer"
	TopUp      TransactionTypes = "Top up"
	GameReward TransactionTypes = "Game Reward"
)

type Transaction struct {
	gorm.Model
	WalletId        uint             `gorm:"not null" json:"wallet_id"`
	TransactionType TransactionTypes `gorm:"type:transaction_types;not null"`
	SourceOfFund    SourceOfFunds    `gorm:"type:source_of_funds"`
	RecipientId     uint             `json:"recipient_id"`
	Amount          decimal.Decimal  `gorm:"not null"`
	Description     string
	Wallet          Wallet `gorm:"foreignKey:wallet_id;references:id"`
	Recipient       Wallet `gorm:"foreignKey:recipient_id;references:id"`
}
