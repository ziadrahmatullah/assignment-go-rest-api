package handler

import "git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/usecase"

type GameHandler struct{
	gu usecase.GameUsecase
}

func NewGameHandler(gu usecase.GameUsecase) *GameHandler{
	return &GameHandler{
		gu: gu,
	}
}
