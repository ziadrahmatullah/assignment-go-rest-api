package dto

import (
	"fmt"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/model"
	"github.com/shopspring/decimal"
)

type ListTransactionsReq struct {
	Search          *string
	FilterStart     *string
	FilterEnd       *string
	SortBy          *string
	SortType        *string
	PaginationLimit *string
	PaginationPage  *string
}

type TopUpReq struct {
	Amount       decimal.Decimal `binding:"required,min=50000,max=10000000" json:"amount"`
	SourceOfFund string          `binding:"required" json:"source_of_fund"`
}

type TransferReq struct {
	WalletNumber string          `binding:"required" json:"to"`
	Amount       decimal.Decimal `binding:"required,min=1000,max=50000000" json:"amount"`
	Description  string          `json:"description,omitempty"`
}

func (tr *TopUpReq) ToTransactionModel(wallet *model.Wallet) model.Transaction {
	tx := model.Transaction{
		WalletId:        wallet.ID,
		TransactionType: model.TransactionTypes(model.TopUp),
		SourceOfFund:    new(model.SourceOfFunds),
		Receiver:        wallet.WalletNumber,
		Amount:          tr.Amount,
		Description:     fmt.Sprintf("Top Up from %s", tr.SourceOfFund),
	}
	*tx.SourceOfFund = model.SourceOfFunds(tr.SourceOfFund)
	return tx
}

func (tr *TransferReq) ToTransactionModel(wallet *model.Wallet) model.Transaction {
	return model.Transaction{
		WalletId:        wallet.ID,
		TransactionType: model.TransactionTypes(model.Transfer),
		Sender:          wallet.WalletNumber,
		Receiver:        tr.WalletNumber,
		Amount:          tr.Amount,
		Description:     tr.Description,
	}
}
