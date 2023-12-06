package usecase_test

import (
	"net/http/httptest"
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/mocks"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/usecase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)
var boxes = []dto.GameBoxesRes{
	{
		ID: 1,
	},
	{
		ID: 2,
	},
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
