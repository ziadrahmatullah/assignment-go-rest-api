package usecase

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/repository"
)

type TransactionUsecase interface {
	GetTransactions(context.Context, dto.ListTransactionsReq) ([]model.Transaction, error)
	TopUp(context.Context, model.Transaction) (*model.Transaction, error)
	Transfer(context.Context, model.Transaction) (*model.Transaction, error)
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

func (tu *transactionUsecase) TopUp(ctx context.Context, req model.Transaction) (transaction *model.Transaction, err error) {
	_, err = tu.wr.FindWallet(ctx, req.WalletId, req.Receiver)
	if err != nil {
		return nil, err
	}
	return tu.tr.TopUpTransaction(ctx, req)
}

func (tu *transactionUsecase) Transfer(ctx context.Context, req model.Transaction) (transaction *model.Transaction, err error) {
	_, err = tu.wr.FindWallet(ctx, req.WalletId, req.Sender)
	if err != nil {
		return nil, err
	}
	return tu.tr.TransferTransaction(ctx, req)
}
