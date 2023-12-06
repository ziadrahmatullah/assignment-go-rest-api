package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/usecase"
	"github.com/gin-gonic/gin"
)

type ResetPasswordHandler struct {
	ru usecase.ResetPasswordUsecase
}

func NewResetPassTokenHandler(ru usecase.ResetPasswordUsecase) *ResetPasswordHandler {
	return &ResetPasswordHandler{
		ru: ru,
	}
}

func (h *ResetPasswordHandler) HandleRequestPassReset(ctx *gin.Context) {
	resp := dto.Response{}
	var req dto.RequestResetPassReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(apperror.ErrInvalidBody)
		return
	}
	res, err := h.ru.RequestPassReset(ctx, req)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.Data = res
	ctx.JSON(http.StatusOK, resp)
}

func (h *ResetPasswordHandler) HandleApplyPassReset(ctx *gin.Context) {
	resp := dto.Response{}
	req := dto.ApplyResetPassReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(apperror.ErrInvalidBody)
		return
	}
	err = h.ru.ApplyPassReset(ctx, req)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.Message = "password has changed"
	ctx.JSON(http.StatusOK, resp)
}
