package middlewares

import (
	"net/http"

	"github.com/ErmoGarcia/mus-backend/utils/token"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := token.TokenValid(token.ExtractToken(c))
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
