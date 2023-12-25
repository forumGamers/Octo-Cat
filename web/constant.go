package web

import "github.com/gin-gonic/gin"

type ResponseWriter interface {
	WriteResponse(c *gin.Context, response WebResponse)
	AbortHttp(c *gin.Context, err error)
	CustomMsgAbortHttp(c *gin.Context, message string, code int)
	HttpValidationErr(c *gin.Context, err error)
	New404Error(msg string) error
	Write200Response(c *gin.Context, msg string, data any)
	Write201Response(c *gin.Context, msg string, data any)
	New403Error(msg string) error
	New401Error(msg string) error
	NewInvalidObjectIdError() error
	New400Error(msg string) error
	New409Error(msg string) error
	New501Error(msg string) error
}

type ResponseWriterImpl struct{}

type RequestReader interface {
	GetParams(c *gin.Context, p any) error
}

type RequestReaderImpl struct{}
