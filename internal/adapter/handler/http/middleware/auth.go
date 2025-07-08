package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"twitter-challenge-exercise/pkg"
)

const (
	authorizationHeader = "Authorization"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader(authorizationHeader)
		bearer := strings.Split(header, " ")

		if len(bearer) < 2 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Authorization header format is invalid",
			})

			ctx.Abort()
			return
		}

		claims, err := pkg.VerifyJWTToken(bearer[1])
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"msg": "Could not verify access token",
			})

			ctx.Abort()
			return
		}

		ctx.Set("user_id", claims.UserID)
		ctx.Next()
	}
}
