package util_test

import (
	"testing"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/util"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestRemoveNewLine(t *testing.T) {
	t.Run("should remove new line", func(t *testing.T) {
		str := "Test\n"
		expected := "Test"

		result := util.RemoveNewLine(str)

		assert.Equal(t, expected, result)
	})
}

func TestToDate(t *testing.T){
	t.Run("should return time.time type", func(t *testing.T) {
		str := "2023-03-03"
		
		result := util.ToDate(str)

		assert.NotNil(t, result)
	})
}

func TestGenerateRandomString(t *testing.T){
	t.Run("should return random string", func(t *testing.T) {		
		result := util.GenerateRandomString()

		assert.NotNil(t, result)
	})
}

func TestIsTopUpAmountValid(t *testing.T){
	t.Run("should return true when valid", func(t *testing.T) {
		amount := decimal.NewFromInt(int64(100000))
		expected := true
		
		result := util.IsTopUpAmountValid(amount)

		assert.Equal(t, expected, result)
	})

	t.Run("should return false when not valid", func(t *testing.T) {
		amount := decimal.NewFromInt(int64(100))
		expected := false
		
		result := util.IsTopUpAmountValid(amount)

		assert.Equal(t, expected, result)
	})
}

func TestIsTransferAmountValid(t *testing.T){
	t.Run("should return true when valid", func(t *testing.T) {
		amount := decimal.NewFromInt(int64(100000))
		expected := true
		
		result := util.IsTransferAmountValid(amount)

		assert.Equal(t, expected, result)
	})

	t.Run("should return false when not valid", func(t *testing.T) {
		amount := decimal.NewFromInt(int64(100))
		expected := false
		
		result := util.IsTopUpAmountValid(amount)

		assert.Equal(t, expected, result)
	})
}


