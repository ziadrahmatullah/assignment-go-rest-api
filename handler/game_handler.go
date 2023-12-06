package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/usecase"
	"github.com/gin-gonic/gin"
)

type GameHandler struct {
	gu usecase.GameUsecase
}

func NewGameHandler(gu usecase.GameUsecase) *GameHandler {
	return &GameHandler{
		gu: gu,
	}
}

func (gh *GameHandler) HandleGetAllBoxes(ctx *gin.Context) {
	resp := dto.Response{}
	boxes, err := gh.gu.GetAllBoxes(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.Data = boxes
	ctx.JSON(http.StatusOK, resp)
}

func (gh *GameHandler) HandleGetRemainingAttempt(ctx *gin.Context) {
	resp := dto.Response{}
	reqContext := dto.CreateContext(ctx)
	res, err := gh.gu.GetRemainingAttempt(ctx, reqContext.UserID)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.Data = res
	ctx.JSON(http.StatusOK, resp)
}

func (gh *GameHandler) HandleChooseBox(ctx *gin.Context) {
	resp := dto.Response{}
	var gameReq dto.GameBoxReq
	err := ctx.ShouldBindJSON(&gameReq)
	if err != nil {
		ctx.Error(apperror.ErrInvalidBody)
		return
	}
	reqContext := dto.CreateContext(ctx)
	res, err := gh.gu.ChooseBox(ctx, gameReq, reqContext.UserID)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.Data = res
	ctx.JSON(http.StatusOK, resp)
}
