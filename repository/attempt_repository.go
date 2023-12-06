package repository

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/model"
	"gorm.io/gorm"
)

type AttemptRepository interface {
	NewAttempt(context.Context, uint) (*model.Attempt, error)
	FindAttempt(context.Context, model.Wallet) (*dto.AttemptRes, error)
}

type attemptRepository struct {
	db *gorm.DB
}

func NewAttemptRepository(db *gorm.DB) AttemptRepository {
	return &attemptRepository{
		db: db,
	}
}

func (ar *attemptRepository) NewAttempt(ctx context.Context, walletId uint) (*model.Attempt, error) {
	newAttempt := &model.Attempt{
		WalletId:         walletId,
		RemainingAttempt: 0,
	}
	err := ar.db.WithContext(ctx).Table("attempts").Create(newAttempt).Error
	if err != nil {
		return nil, err
	}
	return newAttempt, nil
}

func (ar *attemptRepository) FindAttempt(ctx context.Context, wallet model.Wallet) (attempt *dto.AttemptRes, err error) {
	err = ar.db.WithContext(ctx).Table("attempts").Where("wallet_id = ?", wallet.ID).Find(&attempt).Error
	if err != nil {
		return nil, apperror.ErrFindAttemptQuery
	}
	return attempt, nil
}
