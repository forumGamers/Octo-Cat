package middlewares

import (
	"github.com/forumGamers/Octo-Cat/web"
	"github.com/gin-gonic/gin"
)

type Middleware interface {
	Authentication(c *gin.Context)
	Cors() gin.HandlerFunc
	SetContexts(c *gin.Context)
	CheckOrigin(c *gin.Context)
	SetMaxBody(c *gin.Context)
}

type MiddlewareImpl struct {
	web.ResponseWriter
	web.RequestReader
}
