package dto

import (
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/util"
)

type RegisterReq struct {
	Name      string `binding:"required" json:"name"`
	Birthdate string `binding:"required" json:"birtdate"`
	Email     string `binding:"required" json:"email"`
	Password  string `binding:"required" json:"password"`
}

type RegisterRes struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Birthdate string `json:"birtdate"`
	Email     string `json:"email"`
}

type LoginReq struct {
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

type LoginRes struct {
	AccessToken string `json:"accessToken"`
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
