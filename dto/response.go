package dto

import "github.com/gin-gonic/gin"

type Response struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
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
