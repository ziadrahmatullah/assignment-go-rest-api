package apperror

import (
	"net/http"
)

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorRes struct {
	Message string `json:"message"`
}

func NewCustomError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func (ce *CustomError) Error() string {
	return ce.Message
}

func (ce *CustomError) ToErrorRes() ErrorRes {
	return ErrorRes{
		Message: ce.Message,
	}
}

var (
	ErrFindUsersQuery    = NewCustomError(http.StatusInternalServerError, "find user query error")
	ErrFindUserByIdQuery = NewCustomError(http.StatusInternalServerError, "find user by id query error")
	ErrFindUserByEmail   = NewCustomError(http.StatusInternalServerError, "find user by email query error")
	ErrNewUserQuery      = NewCustomError(http.StatusInternalServerError, "new user query error")

	ErrUserNotFound           = NewCustomError(http.StatusBadRequest, "user not found")
	ErrEmailALreadyUsed       = NewCustomError(http.StatusBadRequest, "email already used")
	ErrInvalidPasswordOrEmail = NewCustomError(http.StatusBadRequest, "invalid password or email")

	ErrFindWalletByIdQuery = NewCustomError(http.StatusInternalServerError, "find wallet by id query error")

	ErrWalletNotFound = NewCustomError(http.StatusBadRequest, "wallet not found")

	ErrFindListTransactionQuery  = NewCustomError(http.StatusInternalServerError, "find list transaction query error")
	ErrSortByTransactionQuery    = NewCustomError(http.StatusBadRequest, "wrong key for sorting")
	ErrSortTypeTrasacntionQueqry = NewCustomError(http.StatusBadRequest, "wrong sort type for sorting")
	ErrWrongStartDateFormat      = NewCustomError(http.StatusBadRequest, "wrong start date format")
	ErrWrongEndDateFormat        = NewCustomError(http.StatusBadRequest, "wrong end date format")

	ErrGenerateHashPassword = NewCustomError(http.StatusInternalServerError, "couldn't generate hash password")
	ErrGenerateJWTToken     = NewCustomError(http.StatusInternalServerError, "can't generate jwt token")

	ErrTxCommit = NewCustomError(http.StatusInternalServerError, "commit transaction error")

	ErrInvalidBody = NewCustomError(http.StatusBadRequest, "invalid body")
	ErrUnauthorize = NewCustomError(http.StatusUnauthorized, "unauthorized")
)
