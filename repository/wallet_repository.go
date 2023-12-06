package repository

import (
	"context"
	"fmt"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/model"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type WalletRepository interface {
	FindWalletByUserId(context.Context, uint) (*model.Wallet, error)
	FindWalletByWalletNumber(context.Context, string) (*model.Wallet, error)
	NewWallet(context.Context, uint) (*model.Wallet, error)
}

type walletRepository struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
	return &walletRepository{
		db: db,
	}
}

func (wr *walletRepository) FindWalletByUserId(ctx context.Context, id uint) (wallet *model.Wallet, err error) {
	result := wr.db.WithContext(ctx).Table("wallets").Where("user_id = ?", id).Find(&wallet)
	if result.Error != nil {
		return nil, apperror.ErrFindWalletByIdQuery
	}
	if result.RowsAffected == 0 {
		return nil, apperror.ErrWalletNotFound
	}
	return wallet, nil
}

func (wr *walletRepository) FindWalletByWalletNumber(ctx context.Context, walletNumber string) (wallet *model.Wallet, err error) {
	result := wr.db.WithContext(ctx).Table("wallets").Where("wallet_number= ?", walletNumber).Find(&wallet)
	if result.Error != nil {
		return nil, apperror.ErrFindWalletByIdQuery
	}
	if result.RowsAffected == 0 {
		return nil, apperror.ErrWalletNotFound
	}
	return wallet, nil
}

func (wr *walletRepository) NewWallet(ctx context.Context, userId uint) (*model.Wallet, error) {
	firstThreeDigits := "700"
	var lastWallet model.Wallet
	result := wr.db.WithContext(ctx).Order("id desc").Limit(1).First(&lastWallet)
	if result.Error != nil {
		return nil, apperror.ErrNewWalletQuery
	}
	var nextAutoIncrement int
	if result.RowsAffected > 0 {
		nextAutoIncrement = int(lastWallet.ID) + 1
	} else {
		nextAutoIncrement = 1
	}

	walletNumber := fmt.Sprintf("%s%010d", firstThreeDigits, nextAutoIncrement)

	newWallet := &model.Wallet{
		WalletNumber: walletNumber,
		Balance:      decimal.NewFromInt(int64(0)),
		UserId:       userId,
	}
	err := wr.db.WithContext(ctx).Table("wallets").Create(newWallet).Error
	if err != nil {
		return nil, apperror.ErrNewWalletQuery
	}
	return newWallet, nil
}
