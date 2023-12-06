package usecase

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/repository"
)

type GameUsecase interface {
	GetAllBoxes(context.Context) ([]dto.GameBoxesRes, error)
	GetRemainingAttempt(context.Context, uint) (*dto.AttemptRes, error)
	ChooseBox(context.Context, dto.GameBoxReq, uint) (*dto.ChooseBoxRes, error)
}

type gameUsecase struct {
	gr repository.GameRepository
	wr repository.WalletRepository
	ar repository.AttemptRepository
}

func NewGameUsecase(gr repository.GameRepository, wr repository.WalletRepository, ar repository.AttemptRepository) GameUsecase {
	return &gameUsecase{
		gr: gr,
		wr: wr,
		ar: ar,
	}
}

func (gu *gameUsecase) GetAllBoxes(ctx context.Context) ([]dto.GameBoxesRes, error) {
	return gu.gr.FindAllBoxes(ctx)
}

func (gu *gameUsecase) GetRemainingAttempt(ctx context.Context, userId uint) (*dto.AttemptRes, error) {
	wallet, err := gu.wr.FindWalletByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	return gu.ar.FindAttempt(ctx, *wallet)
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
	attempt, _ := gu.ar.FindAttempt(ctx, *wallet)
	if attempt.RemainingAttempt == 0{
		return nil, apperror.ErrNoAttemptLeft
	}
	return gu.gr.ChooseBox(ctx, *box, *wallet)
}
