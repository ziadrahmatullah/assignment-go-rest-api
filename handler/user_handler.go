package handler

import (
	"net/http"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/usecase"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(uu usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: uu,
	}
}

func (h *UserHandler) HandleGetUsers(ctx *gin.Context) {
	resp := dto.Response{}
	users, err := h.userUsecase.GetAllUsers(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.Data = users
	ctx.JSON(http.StatusOK, resp)
}

func (h *UserHandler) HandleGetUserDetails(ctx *gin.Context) {
	resp := dto.Response{}
	reqContext := dto.CreateContext(ctx)
	user, err := h.userUsecase.GetUserDetails(ctx, reqContext.UserID)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.Data = user
	ctx.JSON(http.StatusOK, resp)
}

func (h *UserHandler) HandleUserRegister(ctx *gin.Context) {
	resp := dto.Response{}
	var data dto.RegisterReq
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.Error(apperror.ErrInvalidBody)
		return
	}
	user, err := h.userUsecase.CreateUser(ctx, data)
	if err != nil {
		ctx.Error(err)
		return
	}
	resp.Data = user
	ctx.JSON(http.StatusOK, resp)

}

func (h *UserHandler) HandleUserLogin(ctx *gin.Context) {
	var data dto.LoginReq
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.Error(apperror.ErrInvalidBody)
		return
	}
	resp, err := h.userUsecase.UserLogin(ctx, data)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
