package repository

import (
	"context"
	"math/rand"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GameRepository interface {
	FindAllBoxes(context.Context) ([]dto.GameBoxesRes, error)
	FindBoxById(context.Context, uint) (*model.Box, error)
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
	rand.Seed(time.Now().UnixNano())
	n := len(boxes)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		boxes[i], boxes[j] = boxes[j], boxes[i]
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

func (gr *gameRepository) ChooseBox(ctx context.Context, box model.Box, wallet model.Wallet) (*dto.ChooseBoxRes, error) {
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
		SourceOfFund:    new(model.SourceOfFunds),
		Receiver:        wallet.WalletNumber,
		Amount:          box.RewardAmount,
	}
	*gameTransaction.SourceOfFund = model.Reward
	tx.Table("transactions").Create(&gameTransaction)
	err := tx.Commit().Error
	if err != nil {
		return nil,err
	}
	var ChoosenBox = dto.ChooseBoxRes{
		RewardAmount: gameTransaction.Amount,
	}
	return &ChoosenBox, nil

}
