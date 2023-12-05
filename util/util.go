package util

import (
	"strings"
	"time"
)

func ToDate(dateString string) time.Time {
	parsedDate, _ := time.Parse("2006-01-02", dateString)
	return parsedDate
}

func RemoveNewLine(str string) string {
	return strings.Trim(str, "\n")
}
