package middlewares

import "github.com/gin-gonic/gin"

func (m *MiddlewareImpl) CheckFileLength(max int, fName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.Request.ParseMultipartForm(int64(100 * 1024 * 1024)); err != nil {
			m.AbortHttp(c, m.New501Error("Failed to parse form data"))
			return
		}

		files := c.Request.MultipartForm.File[fName]
		if files != nil && len(files) > max {
			m.AbortHttp(c, m.New413Error("Maximum allowed files exceeded"))
			return
		}
		m.Next(c)
	}
}
