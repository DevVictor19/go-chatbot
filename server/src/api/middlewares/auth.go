package middlewares

import (
	"net/http"
	"server/src/api/services"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Authorization header is required"})
			ctx.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bearer token is required"})
			ctx.Abort()
			return
		}

		token, err := services.ValidateJwtToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token claims"})
			ctx.Abort()
			return
		}

		email, exists := claims["email"].(string)
		if !exists {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Email field is missing in token"})
			ctx.Abort()
			return
		}

		ctx.Set("email", email)

		ctx.Next()
	}
}
