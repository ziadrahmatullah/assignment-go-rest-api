package dto

import (
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/util"
	"github.com/shopspring/decimal"
)

type RegisterReq struct {
	Name      string `binding:"required" json:"name"`
	Birthdate string `binding:"required" json:"birthdate"`
	Email     string `binding:"required" json:"email"`
	Password  string `binding:"required" json:"password"`
}

type RegisterRes struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Birthdate string `json:"birthdate"`
	Email     string `json:"email"`
}

type LoginReq struct {
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

type LoginRes struct {
	AccessToken string `json:"accessToken"`
}

type UserDetails struct {
	Name         string          `json:"name"`
	Birthdate    time.Time       `json:"birthdate"`
	Email        string          `json:"email"`
	WalletNumber string          `json:"wallet_number"`
	Balance      decimal.Decimal `json:"balance"`
}

func (r *RegisterReq) ToUserModelRegister(password string) model.User {
	return model.User{
		Name:      r.Name,
		Birthdate: util.ToDate(r.Birthdate),
		Email:     r.Email,
		Password:  password,
	}
}

func ToRegisterRes(user model.User) *RegisterRes {
	return &RegisterRes{
		ID:        user.ID,
		Name:      user.Name,
		Birthdate: user.Birthdate.String(),
		Email:     user.Email,
	}
}
