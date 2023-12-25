package middlewares

import (
	"os"

	"github.com/gin-gonic/gin"
)

func (m *MiddlewareImpl) SetContexts(c *gin.Context) {
	stage := os.Getenv("APP_STAGE")
	if stage == "" {
		stage = "Development"
	}

	c.Set("stage", stage)
	m.SetContext(c, "stage", stage)
	m.Next(c)
}
