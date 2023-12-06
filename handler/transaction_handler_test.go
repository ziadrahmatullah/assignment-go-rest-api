package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/server"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/util"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var transactions = []model.Transaction{
	{
		WalletId:        1,
		TransactionType: "Transfer",
		Sender:          "7000000000001",
		Receiver:        "7000000000002",
		Amount:          decimal.NewFromInt(int64(10000)),
	},
}

var transactionsRes = dto.TransactionPaginationRes{
	Data:      transactions,
	TotalData: 20,
	TotalPage: 2,
	Page:      1,
}

var topUpReq = dto.TopUpReq{
	Amount:       decimal.NewFromInt(int64(100000)),
	SourceOfFund: "Cash",
}

var invTopUpReq = dto.TopUpReq{
	Amount: decimal.NewFromInt(int64(100000)),
}

var invTopUpReq2 = dto.TopUpReq{
	Amount:       decimal.NewFromInt(int64(10)),
	SourceOfFund: "Cash",
}

var invTopUpReq3 = dto.TopUpReq{
	Amount:       decimal.NewFromInt(int64(100000)),
	SourceOfFund: "Cassssh",
}

var transferReq = dto.TransferReq{
	WalletNumber: "7000000000001",
	Amount:       decimal.NewFromInt(int64(100000)),
}

var invTransferReq = dto.TransferReq{
	WalletNumber: "7000000000001",
}

var invTransferReq2 = dto.TransferReq{}

func TestHandleGetTransactions(t *testing.T) {
	t.Run("should return 200 if get Transactions success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(transactionsRes)
		tu := mocks.NewTransactionUsecase(t)
		th := handler.NewTransactionHandler(tu)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodGet, "/transactions", nil)
		tu.On("GetTransactions", c, mock.AnythingOfType("dto.ListTransactionsReq")).Return(&transactionsRes, nil)

		th.HandleGetTransactions(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 500 while error in query", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusInternalServerError, "db error")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		tu := mocks.NewTransactionUsecase(t)
		th := handler.NewTransactionHandler(tu)
		tu.On("GetTransactions", mock.Anything, mock.AnythingOfType("dto.ListTransactionsReq")).Return(nil, expectedErr)
		opts := server.RouterOpts{
			TransactionHandler: th,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, "/transactions", nil)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})
}

func TestHandleTopUp(t *testing.T) {
	t.Run("should return 200 if topup success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: transactions[0],
		})
		param, _ := json.Marshal(topUpReq)
		tu := mocks.NewTransactionUsecase(t)
		th := handler.NewTransactionHandler(tu)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodPost, "/transactions/top-up", strings.NewReader(string(param)))
		tu.On("TopUp", c, topUpReq, mock.Anything).Return(&transactions[0], nil)

		th.HandleTopUp(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 400 when invalid body", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusBadRequest, "invalid body")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(invTopUpReq)
		tu := mocks.NewTransactionUsecase(t)
		th := handler.NewTransactionHandler(tu)
		opts := server.RouterOpts{
			TransactionHandler: th,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/transactions/top-up", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 400 when invalid amount", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusBadRequest, "invalid amount")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(invTopUpReq2)
		tu := mocks.NewTransactionUsecase(t)
		th := handler.NewTransactionHandler(tu)
		opts := server.RouterOpts{
			TransactionHandler: th,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/transactions/top-up", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 400 when invalid source of fund", func(t *testing.T) {
		expectedErr := apperror.ErrInvalidSourceOfFund
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(invTopUpReq3)
		tu := mocks.NewTransactionUsecase(t)
		th := handler.NewTransactionHandler(tu)
		opts := server.RouterOpts{
			TransactionHandler: th,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/transactions/top-up", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 500 when error in query", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusInternalServerError, "db error")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(topUpReq)
		tu := mocks.NewTransactionUsecase(t)
		th := handler.NewTransactionHandler(tu)
		tu.On("TopUp", mock.Anything, topUpReq, mock.Anything).Return(nil, expectedErr)
		opts := server.RouterOpts{
			TransactionHandler: th,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/transactions/top-up", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})
}

func TestHandleTransfer(t *testing.T) {
	t.Run("should return 200 if transfer success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: transactions[0],
		})
		param, _ := json.Marshal(transferReq)
		tu := mocks.NewTransactionUsecase(t)
		th := handler.NewTransactionHandler(tu)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodPost, "/transactions/transfer", strings.NewReader(string(param)))
		tu.On("Transfer", c, transferReq, mock.Anything).Return(&transactions[0], nil)

		th.HandleTransfer(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 400 when invalid amount", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusBadRequest, "invalid amount")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(invTransferReq)
		tu := mocks.NewTransactionUsecase(t)
		th := handler.NewTransactionHandler(tu)
		opts := server.RouterOpts{
			TransactionHandler: th,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/transactions/transfer", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 400 when invalid body", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusBadRequest, "invalid body")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(invTransferReq2)
		tu := mocks.NewTransactionUsecase(t)
		th := handler.NewTransactionHandler(tu)
		opts := server.RouterOpts{
			TransactionHandler: th,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/transactions/transfer", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 500 when error in query", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusInternalServerError, "db error")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(transferReq)
		tu := mocks.NewTransactionUsecase(t)
		th := handler.NewTransactionHandler(tu)
		tu.On("Transfer", mock.Anything, transferReq, mock.Anything).Return(nil, expectedErr)
		opts := server.RouterOpts{
			TransactionHandler: th,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/transactions/transfer", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})
}
