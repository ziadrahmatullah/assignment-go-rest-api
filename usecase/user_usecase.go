package usecase

import (
	"context"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/model"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	GetUserDetails(context.Context, uint) (*dto.UserDetails, error)
	GetAllUsers(context.Context) ([]model.User, error)
	CreateUser(context.Context, dto.RegisterReq) (*dto.RegisterRes, error)
	UserLogin(context.Context, dto.LoginReq) (*dto.LoginRes, error)
}

type userUsecase struct {
	ur repository.UserRepository
	wr repository.WalletRepository
	ar repository.AttemptRepository
}

func NewUserUsecase(u repository.UserRepository, wr repository.WalletRepository, ar repository.AttemptRepository) UserUsecase {
	return &userUsecase{
		ur: u,
		wr: wr,
		ar: ar,
	}
}

func (u *userUsecase) GetUserDetails(ctx context.Context, id uint) (*dto.UserDetails, error) {
	return u.ur.FindUserDetails(ctx, id)
}

func (u *userUsecase) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return u.ur.FindUsers(ctx)
}

func (u *userUsecase) CreateUser(ctx context.Context, registerData dto.RegisterReq) (data *dto.RegisterRes, err error) {
	user, _ := u.ur.FindByEmail(ctx, registerData.Email)
	if user != nil {
		return nil, apperror.ErrEmailALreadyUsed
	}
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(registerData.Password), 10)
	userModel := registerData.ToUserModelRegister(string(hashPassword))
	newUser, err := u.ur.NewUser(ctx, userModel)
	if err != nil {
		return nil, err
	}
	newWallet, err := u.wr.NewWallet(ctx, newUser.ID)
	if err != nil {
		return nil, err
	}
	_, err = u.ar.NewAttempt(ctx, newWallet.ID)
	if err != nil {
		return nil, err
	}
	data = dto.ToRegisterRes(*newUser)
	return data, nil
}

func (u *userUsecase) UserLogin(ctx context.Context, loginData dto.LoginReq) (token *dto.LoginRes, err error) {
	user, err := u.ur.FindByEmail(ctx, loginData.Email)
	if err != nil {
		return nil, apperror.ErrInvalidPasswordOrEmail
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		return nil, apperror.ErrInvalidPasswordOrEmail
	}
	newToken, _ := dto.GenerateJWT(dto.JwtClaims{
		ID: user.ID,
	})
	return &dto.LoginRes{AccessToken: newToken}, nil
}
