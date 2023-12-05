package middleware

import (
	"net/http"
	"strings"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"github.com/gin-gonic/gin"
)

func AuthorizeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if gin.Mode() == gin.DebugMode {
			return
		}

		if ctx.Request.URL.Path == "/users/register" {
			ctx.Next()
			return
		}

		if ctx.Request.URL.Path == "/users/login" {
			ctx.Next()
			return
		}

		var resp dto.Response

		header := ctx.GetHeader("Authorization")
		splittedHeader := strings.Split(header, " ")
		if len(splittedHeader) != 2 {
			resp.Message = apperror.ErrUnauthorize.Error()
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}

		token, err := dto.ValidateJWT(splittedHeader[1])
		if err != nil {
			ctx.Error(err)
			return
		}

		claims, ok := token.Claims.(*dto.JwtClaims)
		if !ok || !token.Valid {
			resp.Message = apperror.ErrUnauthorize.Error()
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
			return
		}

		ctx.Set("context", dto.RequestContext{
			UserID: claims.ID,
		})

		ctx.Next()
	}
}
