package dto

import "github.com/gin-gonic/gin"

type Response struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type TransactionPaginationRes struct {
	Data      any `json:"data,omitempty"`
	TotalData int `json:"total_data"`
	TotalPage int `json:"total_page"`
	Page      int `json:"page"`
}

type RequestContext struct {
	UserID uint
}

func CreateContext(ctx *gin.Context) RequestContext {
	res, ok := ctx.Get("context")
	if !ok {
		return RequestContext{}
	}
	return res.(RequestContext)
}
