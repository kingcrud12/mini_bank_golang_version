package middleware

import (
	"net/http"

	"learn_go_with_api/internal/adapter/web/config"
	"learn_go_with_api/internal/application/port/out"

	"github.com/gin-gonic/gin"
)

const CustomerIDContextKey = "customerID"

func AuthMiddleware(tokenGenerator out.TokenGenerator, jwtConfig config.JWTConfig) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie(jwtConfig.CookieName)
		if err != nil || token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authentication required"})
			return
		}

		customerID, err := tokenGenerator.Validate(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		ctx.Set(CustomerIDContextKey, customerID)
		ctx.Next()
	}
}
