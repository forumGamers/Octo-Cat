package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *MiddlewareImpl) CheckOrigin(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		origin := c.Request.Header.Get("Origin")
		if origin == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
			return
		}
	}
	m.Next(c)
}
