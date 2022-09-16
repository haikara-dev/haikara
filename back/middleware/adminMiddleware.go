package middleware

import (
	"context"
	"github.com/haikara-dev/haikara/ent"
	"github.com/haikara-dev/haikara/ent/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminMiddleware(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		UUID := c.MustGet("UUID").(string)

		resUser, err := client.User.
			Query().
			Where(user.UUID(UUID)).
			Only(context.Background())

		if err != nil && !ent.IsNotFound(err) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		if resUser == nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		if resUser.Role != "admin" {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.Set("IS_ADMIN", true)
		c.Next()
	}
}
