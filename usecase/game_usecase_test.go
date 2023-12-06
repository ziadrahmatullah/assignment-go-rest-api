package usecase_test

import (
	"net/http/httptest"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/usecase"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var boxes = []dto.GameBoxesRes{
	{
		ID: 1,
	},
	{
		ID: 2,
	},
}

var wallet1 = model.Wallet{
	WalletNumber: "7000000000001",
	Balance:      decimal.NewFromInt(int64(100000)),
	UserId:       1,
}

var wallet2 = model.Wallet{
	WalletNumber: "7000000000002",
	Balance:      decimal.NewFromInt(int64(0)),
	UserId:       2,
}

var attemptRes0 = dto.AttemptRes{
	RemainingAttempt: 0,
}

var attemptRes = dto.AttemptRes{
	RemainingAttempt: 1,
}

var gameBoxReq = dto.GameBoxReq{
	BoxId: 1,
}

var box = model.Box{
	RewardAmount: decimal.NewFromInt(int64(100000)),
}

var chooseBoxRes = dto.ChooseBoxRes{
	RewardAmount: decimal.NewFromInt(int64(100000)),
}

func TestGetAllBoxes(t *testing.T) {
	t.Run("should return boxes when success", func(t *testing.T) {
		gr := mocks.NewGameRepository(t)
		wr := mocks.NewWalletRepository(t)
		ar := mocks.NewAttemptRepository(t)
		gu := usecase.NewGameUsecase(gr, wr, ar)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		gr.On("FindAllBoxes", c).Return(boxes, nil)

		resUsers, _ := gu.GetAllBoxes(c)

		assert.Equal(t, boxes, resUsers)
	})
}

func TestGetRemainingAttempt(t *testing.T) {
	t.Run("should return attemp when success", func(t *testing.T) {
		gr := mocks.NewGameRepository(t)
		wr := mocks.NewWalletRepository(t)
		ar := mocks.NewAttemptRepository(t)
		gu := usecase.NewGameUsecase(gr, wr, ar)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		wr.On("FindWalletByUserId", c, mock.Anything).Return(&wallet1, nil)
		ar.On("FindAttempt", c, wallet1).Return(&attemptRes, nil)

		resUsers, _ := gu.GetRemainingAttempt(c, uint(1))

		assert.Equal(t, &attemptRes, resUsers)
	})

	t.Run("should return err when wallet not found", func(t *testing.T) {
		expectedErr := apperror.ErrWalletNotFound
		gr := mocks.NewGameRepository(t)
		wr := mocks.NewWalletRepository(t)
		ar := mocks.NewAttemptRepository(t)
		gu := usecase.NewGameUsecase(gr, wr, ar)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		wr.On("FindWalletByUserId", c, mock.Anything).Return(nil, expectedErr)

		_, err := gu.GetRemainingAttempt(c, uint(1))

		assert.ErrorIs(t, err, expectedErr)
	})
}

func TestChooseBox(t *testing.T) {
	t.Run("should return choosen box res when success", func(t *testing.T) {
		gr := mocks.NewGameRepository(t)
		wr := mocks.NewWalletRepository(t)
		ar := mocks.NewAttemptRepository(t)
		gu := usecase.NewGameUsecase(gr, wr, ar)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		wr.On("FindWalletByUserId", c, mock.Anything).Return(&wallet1, nil)
		gr.On("FindBoxById", c, gameBoxReq.BoxId).Return(&box, nil)
		ar.On("FindAttempt", c, wallet1).Return(&attemptRes, nil)
		gr.On("ChooseBox", c, box, wallet1).Return(&chooseBoxRes, nil)

		resUsers, _ := gu.ChooseBox(c, gameBoxReq, uint(1))

		assert.Equal(t, &chooseBoxRes, resUsers)
	})

	t.Run("should return err when remaining attempt 0", func(t *testing.T) {
		expectedErr := apperror.ErrNoAttemptLeft
		gr := mocks.NewGameRepository(t)
		wr := mocks.NewWalletRepository(t)
		ar := mocks.NewAttemptRepository(t)
		gu := usecase.NewGameUsecase(gr, wr, ar)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		wr.On("FindWalletByUserId", c, mock.Anything).Return(&wallet1, nil)
		gr.On("FindBoxById", c, gameBoxReq.BoxId).Return(&box, nil)
		ar.On("FindAttempt", c, wallet1).Return(&attemptRes0, nil)

		_, err := gu.ChooseBox(c, gameBoxReq, uint(1))

		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("should return err when box not found", func(t *testing.T) {
		expectedErr := apperror.ErrWalletNotFound
		gr := mocks.NewGameRepository(t)
		wr := mocks.NewWalletRepository(t)
		ar := mocks.NewAttemptRepository(t)
		gu := usecase.NewGameUsecase(gr, wr, ar)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		wr.On("FindWalletByUserId", c, mock.Anything).Return(&wallet1, nil)
		gr.On("FindBoxById", c, gameBoxReq.BoxId).Return(nil, expectedErr)

		_, err := gu.ChooseBox(c, gameBoxReq, uint(1))

		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("should return err when wallet not found", func(t *testing.T) {
		expectedErr := apperror.ErrBoxNotFound
		gr := mocks.NewGameRepository(t)
		wr := mocks.NewWalletRepository(t)
		ar := mocks.NewAttemptRepository(t)
		gu := usecase.NewGameUsecase(gr, wr, ar)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		wr.On("FindWalletByUserId", c, mock.Anything).Return(nil, expectedErr)

		_, err := gu.ChooseBox(c, gameBoxReq, uint(1))

		assert.ErrorIs(t, err, expectedErr)
	})
}
