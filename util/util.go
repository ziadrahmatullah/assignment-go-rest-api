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

// func ShuffleBoxes(boxes []model.Box) {
// 	rand.Seed(time.Now().UnixNano())
// 	n := len(boxes)
// 	for i := n - 1; i > 0; i-- {
// 		j := rand.Intn(i + 1)
// 		boxes[i], boxes[j] = boxes[j], boxes[i]
// 	}
// }

// func ShuffleArray(arr []int) {
// 	n := len(arr)
// 	for i := n - 1; i > 0; i-- {
// 		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
// 		if err != nil {
// 			panic(err)
// 		}
// 		idx := int(randomIndex.Int64())
// 		arr[i], arr[idx] = arr[idx], arr[i]
// 	}
// }
