package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type SourceOfFunds string
type TransactionTypes string
var AmountReward decimal.Decimal

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
	TransactionType TransactionTypes `gorm:"type:transaction_types;not null" json:"trasaction_type"`
	SourceOfFund    SourceOfFunds    `gorm:"type:source_of_funds" json:"source_of_fund,omitempty"`
	Sender          string           `json:"sender,omitempty"`
	Receiver        string           `json:"receiver"`
	Amount          decimal.Decimal  `gorm:"not null" json:"amount"`
	Description     string           `json:"description,omitempty"`
}
