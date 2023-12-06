package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/util"
	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	tu usecase.TransactionUsecase
}

func NewTransactionHandler(tu usecase.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{
		tu: tu,
	}
}

func (h *TransactionHandler) HandleGetTransactions(ctx *gin.Context) {
	var req dto.ListTransactionsReq
	req.Search = GetStringQueryParam(ctx, "s")
	req.FilterStart = GetStringQueryParam(ctx, "start")
	req.FilterEnd = GetStringQueryParam(ctx, "end")
	req.SortBy = GetStringQueryParam(ctx, "sortBy")
	req.SortType = GetStringQueryParam(ctx, "sort")
	req.PaginationLimit = GetStringQueryParam(ctx, "limit")
	req.PaginationPage = GetStringQueryParam(ctx, "page")
	transactionsRes, err := h.tu.GetTransactions(ctx, req)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, transactionsRes)
}

func (h *TransactionHandler) HandleTopUp(ctx *gin.Context) {
	resp := dto.Response{}
	topUp := dto.TopUpReq{}
	err := ctx.ShouldBindJSON(&topUp)
	if err != nil {
		ctx.Error(apperror.ErrInvalidBody)
		return
	}
	if !util.IsTopUpAmountValid(topUp.Amount) {
		ctx.Error(apperror.ErrInvalidAmount)
		return
	}
	if !model.IsSourceOfFundValid(topUp.SourceOfFund) {
		ctx.Error(apperror.ErrInvalidSourceOfFund)
		return
	}
	reqContext := dto.CreateContext(ctx)
	transactionRes, err := h.tu.TopUp(ctx, topUp, reqContext.UserID)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.Data = transactionRes
	ctx.JSON(http.StatusOK, resp)
}

func (h *TransactionHandler) HandleTransfer(ctx *gin.Context) {
	resp := dto.Response{}
	transfer := dto.TransferReq{}
	err := ctx.ShouldBindJSON(&transfer)
	if err != nil {
		ctx.Error(apperror.ErrInvalidBody)
		return
	}
	if !util.IsTransferAmountValid(transfer.Amount) {
		ctx.Error(apperror.ErrInvalidAmount)
		return
	}
	reqContext := dto.CreateContext(ctx)
	transactionRes, err := h.tu.Transfer(ctx, transfer, reqContext.UserID)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.Data = transactionRes
	ctx.JSON(http.StatusOK, resp)
}

func GetStringQueryParam(ctx *gin.Context, key string) *string {
	value := ctx.Query(key)
	return &value
}
