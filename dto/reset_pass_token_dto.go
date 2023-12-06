package dto

import (
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/model"
)

type RequestResetPassReq struct {
	Email string `binding:"required" json:"email"`
}

type RequestResetPassRes struct {
	Token string `json:"token"`
}

type ApplyResetPassReq struct {
	Email       string `binding:"required" json:"email"`
	NewPassword string `binding:"required" json:"new_password"`
	Token       string `binding:"required" json:"token"`
}

func ToResetPassTokenModel(token string, id uint) *model.ResetPassToken {
	return &model.ResetPassToken{
		Token:  token,
		Expire: time.Now().Add(1 * time.Minute),
		IsUsed: false,
		UserId: id,
	}
}
