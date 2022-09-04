package middleware

import (
	"context"
	"github.com/cubdesign/dailyfj/libs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationToken := c.GetHeader("Authorization")

		idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))
		if idToken == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Id token not available"})
			return
		}

		//verify token
		token, err := libs.FirebaseAuth.VerifyIDToken(context.Background(), idToken)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
			return
		}

		c.Set("UUID", token.UID)
		c.Next()
	}
}
