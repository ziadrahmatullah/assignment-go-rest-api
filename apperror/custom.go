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

	ErrFindJobsQuery            = NewCustomError(http.StatusInternalServerError, "find jobs query error")
	ErrFindJobByIdQuery         = NewCustomError(http.StatusInternalServerError, "find job by id query error")
	ErrNewJobQuery              = NewCustomError(http.StatusInternalServerError, "new job query error")
	ErrUpdateJobExpireDateQuery = NewCustomError(http.StatusInternalServerError, "set job expire date query error")
	ErrRemoveJobQuery           = NewCustomError(http.StatusInternalServerError, "remove job query error")

	ErrInvalidExpireDate    = NewCustomError(http.StatusBadRequest, "invalid expire date")
	ErrJobNotFound          = NewCustomError(http.StatusBadRequest, "job not found")
	ErrJobQuotaHasFulfilled = NewCustomError(http.StatusBadRequest, "job quota has fulfilled")
	ErrAlreadyApplied       = NewCustomError(http.StatusBadRequest, "already apply this job")

	ErrFindRecordsQuery = NewCustomError(http.StatusInternalServerError, "find records query error")
	ErrFindRecordQuery  = NewCustomError(http.StatusInternalServerError, "find record query error")

	ErrRecordNotFound = NewCustomError(http.StatusBadRequest, "record not found")

	ErrGenerateHashPassword = NewCustomError(http.StatusInternalServerError, "couldn't generate hash password")
	ErrGenerateJWTToken     = NewCustomError(http.StatusInternalServerError, "can't generate jwt token")

	ErrTxCommit = NewCustomError(http.StatusInternalServerError, "commit transaction error")

	ErrInvalidBody = NewCustomError(http.StatusBadRequest, "invalid body")
	ErrUnauthorize = NewCustomError(http.StatusUnauthorized, "unauthorized")
)
