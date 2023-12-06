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
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var users = []model.User{
	{
		Name:      "Alice",
		Email:     "alice@gmail.com",
		Birthdate: util.ToDate("2001-03-03"),
	},
	{
		Name:      "Alice",
		Email:     "alice@gmail.com",
		Birthdate: util.ToDate("2001-03-03"),
		Password:  "$2y$12$6jbGWUrwIZquydHg8t1qJOovhmR0f.4u95xN45wLUW24jlFr7q6AG",
	},
}

var registerReq = []dto.RegisterReq{
	{
		Name:      "Alice",
		Birthdate: "2001-03-03",
		Email:     "alice@gmail.com",
		Password:  "alice123",
	},
	{
		Name:     "Alice",
		Email:    "alice@gmail.com",
		Password: "alice123",
	},
}

var loginReq = []dto.LoginReq{
	{
		Email:    "alice@gmail.com",
		Password: "alice123",
	},
	{
		Email: "alice@gmail.com",
	},
}

var userDetails = dto.UserDetails{
	Name:         "Alice",
	Birthdate:    util.ToDate("2001-03-03"),
	Email:        "alice@gmail.com",
	WalletNumber: "7000000000001",
	Balance:      decimal.NewFromInt(int64(100000)),
}

var newWallet = model.Wallet{
	WalletNumber: "7000000000001",
	Balance:      decimal.NewFromInt(int64(100000)),
	UserId:       1,
}

func TestGetAllUsers(t *testing.T) {
	t.Run("should return users when success", func(t *testing.T) {
		ur := mocks.NewUserRepository(t)
		wr := mocks.NewWalletRepository(t)
		ar := mocks.NewAttemptRepository(t)
		uu := usecase.NewUserUsecase(ur, wr, ar)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		ur.On("FindUsers", c).Return(users, nil)

		resUsers, _ := uu.GetAllUsers(c)

		assert.Equal(t, users, resUsers)
	})
}

func TestGetUserDetails(t *testing.T) {
	t.Run("should return users detail when success", func(t *testing.T) {
		ur := mocks.NewUserRepository(t)
		wr := mocks.NewWalletRepository(t)
		ar := mocks.NewAttemptRepository(t)
		uu := usecase.NewUserUsecase(ur, wr, ar)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		ur.On("FindUserDetails", c, mock.Anything).Return(&userDetails, nil)

		resUsers, _ := uu.GetUserDetails(c, uint(1))

		assert.Equal(t, &userDetails, resUsers)
	})
}

func TestCreateUser(t *testing.T) {
	t.Run("should return email already used when it is", func(t *testing.T) {
		expectedErr := apperror.ErrEmailALreadyUsed
		ur := mocks.NewUserRepository(t)
		wr := mocks.NewWalletRepository(t)
		ar := mocks.NewAttemptRepository(t)
		uu := usecase.NewUserUsecase(ur, wr, ar)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		ur.On("FindByEmail", c, mock.Anything).Return(&users[0], nil)

		_, err := uu.CreateUser(c, registerReq[0])

		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("should return registerRes if success", func(t *testing.T) {
		ur := mocks.NewUserRepository(t)
		wr := mocks.NewWalletRepository(t)
		ar := mocks.NewAttemptRepository(t)
		uu := usecase.NewUserUsecase(ur, wr, ar)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		ur.On("FindByEmail", c, registerReq[0].Email).Return(nil, nil)
		ur.On("NewUser", c, mock.Anything).Return(&users[0], nil)
		wr.On("NewWallet", c, mock.Anything).Return(&newWallet, nil)
		ar.On("NewAttempt", c, mock.Anything).Return(nil, nil)

		resUser, _ := uu.CreateUser(c, registerReq[0])

		assert.NotNil(t, resUser)
	})

	t.Run("should return err if new wallet query error", func(t *testing.T) {
		expectedErr := apperror.ErrNewWalletQuery
		ur := mocks.NewUserRepository(t)
		wr := mocks.NewWalletRepository(t)
		ar := mocks.NewAttemptRepository(t)
		uu := usecase.NewUserUsecase(ur, wr, ar)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		ur.On("FindByEmail", c, registerReq[0].Email).Return(nil, nil)
		ur.On("NewUser", c, mock.Anything).Return(&users[0], nil)
		wr.On("NewWallet", c, mock.Anything).Return(nil, expectedErr)

		_, err := uu.CreateUser(c, registerReq[0])

		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("should return err if new attemp query error", func(t *testing.T) {
		expectedErr := apperror.ErrNewAttemptQuery
		ur := mocks.NewUserRepository(t)
		wr := mocks.NewWalletRepository(t)
		ar := mocks.NewAttemptRepository(t)
		uu := usecase.NewUserUsecase(ur, wr, ar)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		ur.On("FindByEmail", c, registerReq[0].Email).Return(nil, nil)
		ur.On("NewUser", c, mock.Anything).Return(&users[0], nil)
		wr.On("NewWallet", c, mock.Anything).Return(&wallet1, nil)
		ar.On("NewAttempt", c, mock.Anything).Return(nil, expectedErr)

		_, err := uu.CreateUser(c, registerReq[0])

		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("should return err if db error", func(t *testing.T) {
		expectedErr := apperror.ErrNewUserQuery
		ur := mocks.NewUserRepository(t)
		wr := mocks.NewWalletRepository(t)
		ar := mocks.NewAttemptRepository(t)
		uu := usecase.NewUserUsecase(ur, wr, ar)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		ur.On("FindByEmail", c, mock.Anything).Return(nil, nil)
		ur.On("NewUser", c, mock.Anything).Return(nil, expectedErr)

		_, err := uu.CreateUser(c, registerReq[0])

		assert.ErrorIs(t, err, expectedErr)
	})
}

func TestLoginUser(t *testing.T) {
	t.Run("should return invalid email or password when it is", func(t *testing.T) {
		expectedErr := apperror.ErrInvalidPasswordOrEmail
		ur := mocks.NewUserRepository(t)
		wr := mocks.NewWalletRepository(t)
		ar := mocks.NewAttemptRepository(t)
		uu := usecase.NewUserUsecase(ur, wr, ar)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		ur.On("FindByEmail", c, mock.Anything).Return(nil, expectedErr)

		_, err := uu.UserLogin(c, loginReq[0])

		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("should return token if success", func(t *testing.T) {
		ur := mocks.NewUserRepository(t)
		wr := mocks.NewWalletRepository(t)
		ar := mocks.NewAttemptRepository(t)
		uu := usecase.NewUserUsecase(ur, wr, ar)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		ur.On("FindByEmail", c, mock.Anything).Return(&users[1], nil)

		token, _ := uu.UserLogin(c, loginReq[0])

		assert.NotNil(t, token)
	})

	t.Run("should return invalid password error", func(t *testing.T) {
		expectedErr := apperror.ErrInvalidPasswordOrEmail
		ur := mocks.NewUserRepository(t)
		wr := mocks.NewWalletRepository(t)
		ar := mocks.NewAttemptRepository(t)
		uu := usecase.NewUserUsecase(ur, wr, ar)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		ur.On("FindByEmail", c, mock.Anything).Return(&users[0], nil)

		_, err := uu.UserLogin(c, loginReq[0])

		assert.ErrorIs(t, err, expectedErr)
	})
}
