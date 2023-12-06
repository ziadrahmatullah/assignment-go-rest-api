package usecase_test

import (
	"net/http/httptest"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/usecase"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var listTransactionsReq = dto.ListTransactionsReq{
	Search:          nil,
	FilterStart:     nil,
	FilterEnd:       nil,
	SortBy:          nil,
	SortType:        nil,
	PaginationLimit: nil,
	PaginationPage:  nil,
}

var transactions = []model.Transaction{
	{
		WalletId:        1,
		TransactionType: "Transfer",
		Sender:          "7000000000001",
		Receiver:        "7000000000002",
		Amount:          decimal.NewFromInt(int64(10000)),
	},
	{
		WalletId:        1,
		TransactionType: "Top Up",
		SourceOfFund:    new(model.SourceOfFunds),
		Receiver:        "7000000000001",
		Amount:          decimal.NewFromInt(int64(10000)),
	},
}

var topUpReq = dto.TopUpReq{
	Amount:       decimal.NewFromInt(int64(100000)),
	SourceOfFund: "Cash",
}

var transferReq = dto.TransferReq{
	WalletNumber: "7000000000001",
	Amount:       decimal.NewFromInt(int64(10000)),
}

var transactionsRes = dto.TransactionPaginationRes{
	Data:      transactions,
	TotalData: 20,
	TotalPage: 2,
	Page:      1,
}

func TestGetTransactions(t *testing.T) {
	t.Run("should return transaction pagination res when success", func(t *testing.T) {
		tr := mocks.NewTransactionRepository(t)
		wr := mocks.NewWalletRepository(t)
		tu := usecase.NewTransactionUsecase(tr, wr)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		tr.On("FindListTransaction", c, mock.Anything).Return(&transactionsRes, nil)
		expected := transactionsRes

		resUsers, _ := tu.GetTransactions(c, listTransactionsReq)

		assert.Equal(t, &expected, resUsers)
	})
}

func TestTopUp(t *testing.T) {
	t.Run("should return transaction when success", func(t *testing.T) {
		tr := mocks.NewTransactionRepository(t)
		wr := mocks.NewWalletRepository(t)
		tu := usecase.NewTransactionUsecase(tr, wr)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		wr.On("FindWalletByUserId", c, mock.Anything).Return(&wallet1, nil)
		tr.On("TopUpTransaction", c, mock.Anything).Return(&transactions[1], nil)

		resUsers, _ := tu.TopUp(c, topUpReq, uint(1))

		assert.NotNil(t, resUsers)
	})

	t.Run("should return err when wallet not found", func(t *testing.T) {
		expectedErr := apperror.ErrWalletNotFound
		tr := mocks.NewTransactionRepository(t)
		wr := mocks.NewWalletRepository(t)
		tu := usecase.NewTransactionUsecase(tr, wr)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		wr.On("FindWalletByUserId", c, mock.Anything).Return(nil, expectedErr)

		_, err := tu.TopUp(c, topUpReq, uint(1))

		assert.ErrorIs(t, err, expectedErr)
	})
}

func TestTransfer(t *testing.T) {
	t.Run("should return transaction when success", func(t *testing.T) {
		tr := mocks.NewTransactionRepository(t)
		wr := mocks.NewWalletRepository(t)
		tu := usecase.NewTransactionUsecase(tr, wr)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		wr.On("FindWalletByUserId", c, mock.Anything).Return(&wallet1, nil)
		tr.On("TopUpTransaction", c, mock.Anything).Return(&transactions[1], nil)

		resUsers, _ := tu.TopUp(c, topUpReq, uint(1))

		assert.NotNil(t, resUsers)
	})

	t.Run("should return err when wallet not found", func(t *testing.T) {
		expectedErr := apperror.ErrWalletNotFound
		tr := mocks.NewTransactionRepository(t)
		wr := mocks.NewWalletRepository(t)
		tu := usecase.NewTransactionUsecase(tr, wr)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		wr.On("FindWalletByUserId", c, mock.Anything).Return(nil, expectedErr)

		_, err := tu.Transfer(c, transferReq, uint(1))

		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("should return err when wallet number not found", func(t *testing.T) {
		expectedErr := apperror.ErrInvalidWalletNumber
		tr := mocks.NewTransactionRepository(t)
		wr := mocks.NewWalletRepository(t)
		tu := usecase.NewTransactionUsecase(tr, wr)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		wr.On("FindWalletByUserId", c, mock.Anything).Return(&wallet1, nil)
		wr.On("FindWalletByWalletNumber", c, mock.Anything).Return(nil, expectedErr)

		_, err := tu.Transfer(c, transferReq, uint(1))

		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("should return err when wallet sender and receiver is same", func(t *testing.T) {
		expectedErr := apperror.ErrCantTransferToYourWallet
		tr := mocks.NewTransactionRepository(t)
		wr := mocks.NewWalletRepository(t)
		tu := usecase.NewTransactionUsecase(tr, wr)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		wr.On("FindWalletByUserId", c, mock.Anything).Return(&wallet1, nil)
		wr.On("FindWalletByWalletNumber", c, mock.Anything).Return(&wallet1, nil)

		_, err := tu.Transfer(c, transferReq, uint(1))

		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("should return err when insufficient balance", func(t *testing.T) {
		expectedErr := apperror.ErrInsufficientBalance
		tr := mocks.NewTransactionRepository(t)
		wr := mocks.NewWalletRepository(t)
		tu := usecase.NewTransactionUsecase(tr, wr)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		wr.On("FindWalletByUserId", c, mock.Anything).Return(&wallet2, nil)
		wr.On("FindWalletByWalletNumber", c, mock.Anything).Return(&wallet1, nil)

		_, err := tu.Transfer(c, transferReq, uint(1))

		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("should transaction when success", func(t *testing.T) {
		tr := mocks.NewTransactionRepository(t)
		wr := mocks.NewWalletRepository(t)
		tu := usecase.NewTransactionUsecase(tr, wr)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		wr.On("FindWalletByUserId", c, mock.Anything).Return(&wallet1, nil)
		wr.On("FindWalletByWalletNumber", c, mock.Anything).Return(&wallet2, nil)
		tr.On("TransferTransaction", c, transferReq.ToTransactionModel(&wallet1)).Return(&transactions[0], nil)
		expected := transactions[0]

		resUser, _ := tu.Transfer(c, transferReq, uint(1))

		assert.Equal(t, &expected, resUser)
	})
}
