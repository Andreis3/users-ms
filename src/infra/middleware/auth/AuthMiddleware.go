package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	userKey string
}

func NewAuthMiddleware(userKey string) *AuthMiddleware {
	return &AuthMiddleware{
		userKey: userKey,
	}
}
func (m AuthMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.GetHeader("USER-KEY")
		if user != m.userKey {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
