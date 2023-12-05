package repository

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/model"
	"gorm.io/gorm"
)

type WalletRepository interface {
	FindWalletByUserId(context.Context, uint) (*model.Wallet, error)
	FindWalletByWalletNumber(context.Context, string) (*model.Wallet, error)
	FindWallet(context.Context, uint, string) (*model.Wallet, error)
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
	result := wr.db.WithContext(ctx).Table("wallets").Where("wallet_number = ?", walletNumber).Find(&wallet)
	if result.Error != nil {
		return nil, apperror.ErrFindWalletByIdQuery
	}
	if result.RowsAffected == 0 {
		return nil, apperror.ErrWalletNotFound
	}
	return wallet, nil
}

func (wr *walletRepository) FindWallet(ctx context.Context, id uint, walletNumber string) (wallet *model.Wallet, err error) {
	result := wr.db.WithContext(ctx).Table("wallets").Where("wallet_number = ? AND id = ?", walletNumber, id).Find(&wallet)
	if result.Error != nil {
		return nil, apperror.ErrFindWalletByIdQuery
	}
	if result.RowsAffected == 0 {
		return nil, apperror.ErrWalletNotFound
	}
	return wallet, nil
}
