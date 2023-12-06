package usecase

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/repository"
)

type TransactionUsecase interface {
	GetTransactions(context.Context, dto.ListTransactionsReq) (*dto.TransactionPaginationRes, error)
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

func (tu *transactionUsecase) GetTransactions(ctx context.Context, req dto.ListTransactionsReq) (transactions *dto.TransactionPaginationRes, err error) {
	return tu.tr.FindListTransaction(ctx, req)
}

func (tu *transactionUsecase) TopUp(ctx context.Context, req dto.TopUpReq, userId uint) (transaction *model.Transaction, err error) {
	wallet, err := tu.wr.FindWalletByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	newTransaction := req.ToTransactionModel(wallet)
	return tu.tr.TopUpTransaction(ctx, newTransaction)
}

func (tu *transactionUsecase) Transfer(ctx context.Context, req dto.TransferReq, userId uint) (transaction *model.Transaction, err error) {
	senderWallet, err := tu.wr.FindWalletByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	receiverWallet, err := tu.wr.FindWalletByWalletNumber(ctx, req.WalletNumber)
	if err != nil {
		return nil, apperror.ErrInvalidWalletNumber
	}
	if senderWallet.WalletNumber == receiverWallet.WalletNumber {
		return nil, apperror.ErrCantTransferToYourWallet
	}
	if senderWallet.Balance.LessThan(req.Amount) {
		return nil, apperror.ErrInsufficientBalance
	}
	newTransaction := req.ToTransactionModel(senderWallet)
	return tu.tr.TransferTransaction(ctx, newTransaction)
}
