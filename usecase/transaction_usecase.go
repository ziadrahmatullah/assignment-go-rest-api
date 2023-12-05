package usecase

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/repository"
)

type TransactionUsecase interface {
	GetTransactions(context.Context, dto.ListTransactionsReq) ([]model.Transaction, error)
	TopUp(context.Context, dto.TopUpReq, uint) (*model.Transaction, error)
	Transfer(context.Context, dto.TransferReq, uint) (*model.Transaction, error)
}

type transactionUsecase struct {
	tr repository.TransactionRepository
	wr repository.WalletRepository
}

func NewTransactionUsecase(tr repository.TransactionRepository, wr repository.WalletRepository) TransactionUsecase {
	return &transactionUsecase{
		tr: tr,
		wr: wr,
	}
}

func (tu *transactionUsecase) GetTransactions(ctx context.Context, req dto.ListTransactionsReq) (transactions []model.Transaction, err error) {
	return tu.tr.FindListTransaction(ctx, req)
}

func (tu *transactionUsecase) TopUp(ctx context.Context, req dto.TopUpReq, id uint) (transaction *model.Transaction, err error) {
	wallet, err := tu.wr.FindWalletById(ctx, id)
	if err != nil {
		return nil, err
	}
	newTransaction := req.ToTransactionModel(wallet)
	return tu.tr.TopUpTransaction(ctx, newTransaction)
}

func (tu *transactionUsecase) Transfer(ctx context.Context, req dto.TransferReq, id uint) (transaction *model.Transaction, err error) {
	wallet, err := tu.wr.FindWalletById(ctx, id)
	if err != nil {
		return nil, err
	}
	newTransaction := req.ToTransactionModel(wallet)
	return tu.tr.TransferTransaction(ctx, newTransaction)
}
