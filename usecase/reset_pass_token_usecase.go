package usecase

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/repository"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/util"
)

type ResetPassTokenUsecase interface {
	RequestPassReset(context.Context, dto.RequestResetPassReq) (*dto.RequestResetPassRes, error)
	ApplyPassReset(context.Context, dto.ApplyResetPassReq) error
}

type resetPassTokenUsecase struct {
	rr repository.ResetPassTokenRepository
	ur repository.UserRepository
}

func NewResetPassTokenUsecase(rr repository.ResetPassTokenRepository, ur repository.UserRepository) ResetPassTokenUsecase {
	return &resetPassTokenUsecase{
		rr: rr,
		ur: ur,
	}
}

func (ru *resetPassTokenUsecase) RequestPassReset(ctx context.Context, req dto.RequestResetPassReq) (res *dto.RequestResetPassRes, err error) {
	user, err := ru.ur.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, apperror.ErrEmailNotFound
	}
	token := util.GenerateRandomString()
	resetPassTokenModel := dto.ToResetPassTokenModel(token, user.ID)
	err = ru.rr.CreateResetPassToken(ctx, *resetPassTokenModel)
	if err != nil {
		return nil, err
	}
	res.Token = token
	return res, nil
}

func (ru *resetPassTokenUsecase) ApplyPassReset(ctx context.Context, req dto.ApplyResetPassReq) error {
	return ru.rr.ApplyResetPassToken(ctx, req)
}
