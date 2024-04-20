package middleware

import (
	"net/http"
	"strings"

	"github.com/danielsobrado/ainovelprompter/pkg/auth"
	"github.com/danielsobrado/ainovelprompter/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if viper.GetBool("auth.enabled") {
			authHeader := c.GetHeader("Authorization")
			if authHeader == "" {
				logging.Logger.Error("Missing Authorization header")
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}

			bearerToken := strings.Split(authHeader, " ")
			if len(bearerToken) != 2 {
				logging.Logger.Error("Invalid Authorization header format")
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}

			tokenString := bearerToken[1]
			claims, err := auth.ValidateToken(tokenString)
			if err != nil {
				logging.Logger.Errorf("Invalid token: %v", err)
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}

			c.Set("userID", claims.UserID)
			c.Set("role", claims.Role)
		}

		c.Next()
	}
}
