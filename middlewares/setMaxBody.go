package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *MiddlewareImpl) SetMaxBody(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 100<<20)

	m.Next(c)
}
