package usecase_test

import (
	"net/http/httptest"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/usecase"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var requestResetPassReq = dto.RequestResetPassReq{
	Email: "alice@gmail.com",
}

var requestPassResetRes = dto.RequestResetPassRes{
	Token: "example",
}

var user = model.User{
	Name:      "Alice",
	Email:     "alice@gmail.com",
	Birthdate: util.ToDate("2001-03-03"),
}

var applyResetPassReq = dto.ApplyResetPassReq{
	Email: "alice@gmail.com",
	NewPassword: "alice111",
	Token: "example",
}

func TestRequestPassReset(t *testing.T) {
	t.Run("should return request reset pass res when success", func(t *testing.T) {
		rr := mocks.NewResetPassTokenRepository(t)
		ur := mocks.NewUserRepository(t)
		gu := usecase.NewResetPassTokenUsecase(rr, ur)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		ur.On("FindByEmail", c, requestResetPassReq.Email).Return(&user, nil)
		rr.On("CreateResetPassToken", c, mock.Anything).Return(nil)

		resUsers, _ := gu.RequestPassReset(c, requestResetPassReq)

		assert.NotNil(t,resUsers)
	})
	t.Run("should return err when email not found", func(t *testing.T) {
		expectedErr := apperror.ErrEmailNotFound
		rr := mocks.NewResetPassTokenRepository(t)
		ur := mocks.NewUserRepository(t)
		gu := usecase.NewResetPassTokenUsecase(rr, ur)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		ur.On("FindByEmail", c, requestResetPassReq.Email).Return(nil, expectedErr)

		_, err := gu.RequestPassReset(c, requestResetPassReq)

		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("should return err when email not found", func(t *testing.T) {
		expectedErr := apperror.ErrCreateResetPassTokenQuery
		rr := mocks.NewResetPassTokenRepository(t)
		ur := mocks.NewUserRepository(t)
		gu := usecase.NewResetPassTokenUsecase(rr, ur)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		ur.On("FindByEmail", c, requestResetPassReq.Email).Return(&user, nil)
		rr.On("CreateResetPassToken", c, mock.Anything).Return(expectedErr)

		_, err := gu.RequestPassReset(c, requestResetPassReq)

		assert.ErrorIs(t, err, expectedErr)
	})
}

func TestApplyPassReset(t *testing.T) {
	t.Run("should return apply reset pass res when success", func(t *testing.T) {
		rr := mocks.NewResetPassTokenRepository(t)
		ur := mocks.NewUserRepository(t)
		gu := usecase.NewResetPassTokenUsecase(rr, ur)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		rr.On("ApplyResetPassToken", c, mock.Anything).Return(nil)

		err := gu.ApplyPassReset(c, applyResetPassReq)

		assert.Nil(t,err)
	})
}
