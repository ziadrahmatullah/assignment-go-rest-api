package util

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

func ToDate(dateString string) time.Time {
	parsedDate, _ := time.Parse("2006-01-02", dateString)
	return parsedDate
}

func RemoveNewLine(str string) string {
	return strings.Trim(str, "\n")
}

func GenerateRandomString() string {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}

func IsTopUpAmountValid(amount decimal.Decimal) bool {
	minAmount := decimal.NewFromInt(50000)
	maxAmount := decimal.NewFromInt(10000000)

	if amount.LessThan(minAmount) || amount.GreaterThan(maxAmount) {
		return false
	}
	return true
}

func IsTransferAmountValid(amount decimal.Decimal) bool {
	minAmount := decimal.NewFromInt(1000)
	maxAmount := decimal.NewFromInt(50000000)

	if amount.LessThan(minAmount) || amount.GreaterThan(maxAmount) {
		return false
	}
	return true
}
