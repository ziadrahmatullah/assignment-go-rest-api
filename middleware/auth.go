package middleware

import (
	"net/http"
	"strings"
	"time"

	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/apperror"
	"git.garena.com/sea-labs-id/bootcamp/batch-02/ziad-rahmatullah/assignment-go-rest-api/dto"
	"github.com/gin-gonic/gin"
)

func AuthorizeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if gin.Mode() == gin.DebugMode {
			return
		}

		unauthorizedResponse := func() {
			var resp dto.Response
			resp.Message = apperror.ErrUnauthorize.Error()
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, resp)
		}

		excludedPaths := []string{
			"/users/register",
			"/users/login",
			"/users/reset-password",
		}

		for _, path := range excludedPaths {
			if ctx.Request.URL.Path == path {
				ctx.Next()
				return
			}
		}

		header := ctx.GetHeader("Authorization")
		splittedHeader := strings.Split(header, " ")
		if len(splittedHeader) != 2 {
			unauthorizedResponse()
			return
		}

		token, err := dto.ValidateJWT(splittedHeader[1])
		if err != nil {
			ctx.Error(err)
			unauthorizedResponse()
			return
		}

		claims, ok := token.Claims.(*dto.JwtClaims)
		if !ok || !token.Valid || claims.ExpiresAt.Before(time.Now()) {
			unauthorizedResponse()
			return
		}
		ctx.Set("context", dto.RequestContext{
			UserID: claims.ID,
		})

		ctx.Next()
	}
}
