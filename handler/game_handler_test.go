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

var attemptRes = dto.AttemptRes{
	RemainingAttempt: 1,
}

var gameBoxReq = dto.GameBoxReq{
	BoxId: 1,
}

var invGameBoxReq = dto.GameBoxReq{}

var choosenBoxRes = dto.ChooseBoxRes{
	RewardAmount: decimal.NewFromInt(int64(10000)),
}

func TestHandleGetAllBoxes(t *testing.T) {
	t.Run("should return 200 if get all boxes success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: boxes,
		})
		gu := mocks.NewGameUsecase(t)
		gh := handler.NewGameHandler(gu)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodGet, "/games/boxes", nil)
		gu.On("GetAllBoxes", c).Return(boxes, nil)

		gh.HandleGetAllBoxes(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 500 while error in query", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusInternalServerError, "db error")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		gu := mocks.NewGameUsecase(t)
		gh := handler.NewGameHandler(gu)
		gu.On("GetAllBoxes", mock.Anything).Return(nil, expectedErr)
		opts := server.RouterOpts{
			GameHandler: gh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, "/games/boxes", nil)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})
}

func TestHandleGetRemainingAttempt(t *testing.T) {
	t.Run("should return 200 if get remaining attempt success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: attemptRes,
		})
		gu := mocks.NewGameUsecase(t)
		gh := handler.NewGameHandler(gu)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodGet, "/games/attempts", nil)
		gu.On("GetRemainingAttempt", c, mock.Anything).Return(&attemptRes, nil)

		gh.HandleGetRemainingAttempt(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 500 while error in query", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusInternalServerError, "db error")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		gu := mocks.NewGameUsecase(t)
		gh := handler.NewGameHandler(gu)
		gu.On("GetRemainingAttempt", mock.Anything, mock.Anything).Return(nil, expectedErr)
		opts := server.RouterOpts{
			GameHandler: gh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodGet, "/games/attempts", nil)
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})
}

func TestHandleChooseBox(t *testing.T) {
	t.Run("should return 200 if choose box success", func(t *testing.T) {
		expectedResp, _ := json.Marshal(dto.Response{
			Data: choosenBoxRes,
		})
		param, _ := json.Marshal(gameBoxReq)
		gu := mocks.NewGameUsecase(t)
		gh := handler.NewGameHandler(gu)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request, _ = http.NewRequest(http.MethodPost, "/games", strings.NewReader(string(param)))
		gu.On("ChooseBox", c, gameBoxReq, mock.Anything).Return(&choosenBoxRes, nil)

		gh.HandleChooseBox(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResp), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 400 when invalid body", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusBadRequest, "invalid body")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(invGameBoxReq)
		gu := mocks.NewGameUsecase(t)
		gh := handler.NewGameHandler(gu)
		opts := server.RouterOpts{
			GameHandler: gh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/games", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})

	t.Run("should return 500 when error in query", func(t *testing.T) {
		expectedErr := apperror.NewCustomError(http.StatusInternalServerError, "db error")
		resBody, _ := json.Marshal(expectedErr.ToErrorRes())
		param, _ := json.Marshal(gameBoxReq)
		gu := mocks.NewGameUsecase(t)
		gh := handler.NewGameHandler(gu)
		gu.On("ChooseBox", mock.Anything, gameBoxReq, mock.Anything).Return(nil, expectedErr)
		opts := server.RouterOpts{
			GameHandler: gh,
		}
		r := server.NewRouter(opts)
		rec := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/games", strings.NewReader(string(param)))
		r.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(resBody), util.RemoveNewLine(rec.Body.String()))
	})
}
