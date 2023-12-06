package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/handler"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/server"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var requestPassResetReq = dto.RequestResetPassReq{
	Email: "alice@gmail.com",
}

var invRequestPassResetReq = dto.RequestResetPassReq{}

var requestPassResetRes = dto.RequestResetPassRes{
	Token: "example",
}

var applyResetPassReq = dto.ApplyResetPassReq{
	Email:       "alice@gmail.com",
	NewPassword: "alice123",
	Token:       "example",
}

var invApplyResetPassReq = dto.ApplyResetPassReq{
	NewPassword: "alice123",
	Token:       "example",
}

func TestHandleRequestPassReset(t *testing.T) {
	t.Run("should return 200 if requst password reset success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: requestPassResetRes,
		})
		param, _ := json.Marshal(requestPassResetReq)
		ru := mocks.NewResetPasswordUsecase(t)
		rh := handler.NewResetPassTokenHandler(ru)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodPost, "/users/rest-password", strings.NewReader(string(param)))
		ru.On("RequestPassReset", c, requestPassResetReq).Return(&requestPassResetRes, nil)

		rh.HandleRequestPassReset(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 400 when invalid body", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusBadRequest, "invalid body")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(invRequestPassResetReq)
		ru := mocks.NewResetPasswordUsecase(t)
		rh := handler.NewResetPassTokenHandler(ru)
		opts := server.RouterOpts{
			ResetPasswordHandler: rh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/users/reset-password", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 500 when error in query", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusInternalServerError, "db error")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(requestPassResetReq)
		ru := mocks.NewResetPasswordUsecase(t)
		rh := handler.NewResetPassTokenHandler(ru)
		ru.On("RequestPassReset", mock.Anything, requestPassResetReq).Return(nil, expectedErr)
		opts := server.RouterOpts{
			ResetPasswordHandler: rh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/users/reset-password", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})
}

func TestHandleApplyPassReset(t *testing.T) {
	t.Run("should return 200 if topup success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Message: "password has changed",
		})
		param, _ := json.Marshal(applyResetPassReq)
		ru := mocks.NewResetPasswordUsecase(t)
		rh := handler.NewResetPassTokenHandler(ru)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodPut, "/users/rest-password", strings.NewReader(string(param)))
		ru.On("ApplyPassReset", c, applyResetPassReq).Return(nil)

		rh.HandleApplyPassReset(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 400 when invalid body", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusBadRequest, "invalid body")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(invApplyResetPassReq)
		ru := mocks.NewResetPasswordUsecase(t)
		rh := handler.NewResetPassTokenHandler(ru)
		opts := server.RouterOpts{
			ResetPasswordHandler: rh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPut, "/users/reset-password", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 500 when error in query", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusInternalServerError, "db error")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(applyResetPassReq)
		ru := mocks.NewResetPasswordUsecase(t)
		rh := handler.NewResetPassTokenHandler(ru)
		ru.On("ApplyPassReset", mock.Anything, applyResetPassReq).Return(expectedErr)
		opts := server.RouterOpts{
			ResetPasswordHandler: rh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPut, "/users/reset-password", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})
}
