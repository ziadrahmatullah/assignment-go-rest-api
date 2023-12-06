package dto

import "github.com/shopspring/decimal"

type GameBoxReq struct {
	BoxId uint `binding:"required" json:"box_id"`
}

type GameBoxesRes struct {
	ID uint `json:"id"`
}

type AttemptRes struct {
	RemainingAttempt int `json:"remaining_attempt"`
}

type ChooseBoxRes struct {
	RewardAmount decimal.Decimal `json:"reward_amount"`
}
