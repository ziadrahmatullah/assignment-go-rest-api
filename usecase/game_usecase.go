package usecase

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/repository"
)

type GameUsecase interface {
}

type gameUsecase struct {
	gr repository.GameRepository
	wr repository.WalletRepository
}

func NewGameUsecase (gr repository.GameRepository, wr repository.WalletRepository) GameUsecase{
	return &gameUsecase{
		gr: gr,
		wr: wr,
	}
}

func (gu *gameUsecase) GetAllBoxs(ctx context.Context) ([]dto.GameBoxsRes, error) {
	return gu.gr.FindAllBoxs(ctx)
}

func (gu *gameUsecase) GetRemainingAttempt(ctx context.Context, userId uint) (*dto.AttemptRes, error) {
	wallet, err := gu.wr.FindWalletByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	return gu.gr.FindAttempt(ctx, *wallet)
}

func (gu *gameUsecase) ChooseBox(ctx context.Context, req dto.GameBoxReq, userId uint) (*dto.ChooseBoxRes, error) {
	wallet, err := gu.wr.FindWalletByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	box, err := gu.gr.FindBoxById(ctx, req.BoxId)
	if err != nil {
		return nil, err
	}
	return gu.gr.ChooseBox(ctx, *box, *wallet)
}
