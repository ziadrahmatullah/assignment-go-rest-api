package repository

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GameRepository interface {
	FindAllBoxes(context.Context) ([]dto.GameBoxesRes, error)
	FindBoxById(context.Context, uint) (*model.Box, error)
	FindAttempt(context.Context, model.Wallet) (*dto.AttemptRes, error)
	ChooseBox(context.Context, model.Box, model.Wallet) (*dto.ChooseBoxRes, error)
}

type gameRepository struct {
	db *gorm.DB
}

func NewGameRepository(db *gorm.DB) GameRepository {
	return &gameRepository{
		db: db,
	}
}

func (gr *gameRepository) FindAllBoxes(ctx context.Context) (boxes []dto.GameBoxesRes, err error) {
	err = gr.db.WithContext(ctx).Table("boxes").Find(&boxes).Error
	if err != nil {
		return nil, apperror.ErrFindBoxesQuery
	}
	return boxes, nil
}

func (gr *gameRepository) FindBoxById(ctx context.Context, id uint) (box *model.Box, err error) {
	result := gr.db.WithContext(ctx).Table("boxes").Where("id = ?", id).Find(&box)
	if result.Error != nil {
		return nil, apperror.ErrFindBoxByIdQuery
	}
	if result.RowsAffected == 0 {
		return nil, apperror.ErrBoxNotFound
	}
	return box, nil
}

func (gr *gameRepository) FindAttempt(ctx context.Context, wallet model.Wallet) (attempt *dto.AttemptRes, err error) {
	err = gr.db.WithContext(ctx).Table("attempts").Where("wallet_id = ?", wallet.ID).Find(&attempt).Error
	if err != nil {
		return nil, apperror.ErrFindAttemptQuery
	}
	return attempt, nil
}

func (gr *gameRepository) ChooseBox(ctx context.Context, box model.Box, wallet model.Wallet) (ChoosenBox *dto.ChooseBoxRes, err error) {
	tx := gr.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	tx.Table("attempts").
		Where("wallet_id = ?", wallet.ID).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Update("remaining_attempt", gorm.Expr("remaining_attempt - ?", 1))
	tx.Table("wallets").
		Where("id = ?", wallet.ID).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Update("balance", gorm.Expr("balance + ?", box.RewardAmount))

	gameHistory := &model.Game{
		WalletId: wallet.ID,
		BoxId:    box.ID,
	}
	tx.Table("games").Create(&gameHistory)

	gameTransaction := &model.Transaction{
		WalletId:        wallet.ID,
		TransactionType: model.GameReward,
		SourceOfFund:    model.Reward,
		Receiver:        wallet.WalletNumber,
		Amount:          box.RewardAmount,
	}
	tx.Table("transactions").Create(&gameTransaction)
	err = tx.Commit().Error
	if err != nil {
		return nil, apperror.ErrTxCommit
	}
	ChoosenBox.RewardAmount = gameTransaction.Amount
	return ChoosenBox, nil

}
