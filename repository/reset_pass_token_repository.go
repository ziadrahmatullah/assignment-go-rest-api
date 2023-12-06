package repository

import (
	"context"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/model"
	"gorm.io/gorm"
)

type ResetPassTokenRepository interface {
	CreateResetPassToken(context.Context, model.ResetPassToken) error
	ApplyResetPassToken(context.Context, dto.ApplyResetPassReq) error
}

type resetPassTokenRepository struct {
	db *gorm.DB
}

func NewResetPassTokenRepository(db *gorm.DB) ResetPassTokenRepository {
	return &resetPassTokenRepository{
		db: db,
	}
}

func (rr *resetPassTokenRepository) CreateResetPassToken(ctx context.Context, req model.ResetPassToken) (err error) {
	err = rr.db.WithContext(ctx).Table("reset_pass_tokens").Create(&req).Error
	if err != nil {
		return apperror.ErrCreateResetPassTokenQuery
	}
	return nil
}

func (rr *resetPassTokenRepository) ApplyResetPassToken(ctx context.Context, req dto.ApplyResetPassReq) (err error) {
	var resetToken model.ResetPassToken
	result := rr.db.WithContext(ctx).Where("token = ? AND is_used = false", req.Token).Preload("User").First(&resetToken)
	if result.Error != nil {
		return apperror.ErrInvalidToken
	}
	if resetToken.User.Email != req.Email {
		return apperror.ErrInvalidEmail
	}
	if resetToken.Expire.Before(time.Now()) {
		return apperror.ErrTokenExpired
	}
	resetToken.User.Password = req.NewPassword
	rr.db.WithContext(ctx).Save(&resetToken.User)
	resetToken.IsUsed = true
	rr.db.WithContext(ctx).Save(&resetToken)
	return nil
}
