package web

import (
	"github.com/forumGamers/Octo-Cat/errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type WebResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type HttpError struct {
	Error   error
	Code    int
	Message string
}

type InputError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewResponseWriter() ResponseWriter {
	return &ResponseWriterImpl{}
}

func (w *ResponseWriterImpl) getStatusMessage(status int) string {
	statusMessages := map[int]string{
		100: "Continue",
		101: "Switching Protocols",
		102: "Processing",
		103: "Early Hints",
		200: "OK",
		201: "Created",
		202: "Accepted",
		203: "Non Authoritative Info",
		204: "No Content",
		205: "Reset Content",
		206: "Partial Content",
		207: "Multi Status",
		208: "Already Reported",
		226: "IM Used",
		300: "Multiple Choices",
		301: "Moved Permanently",
		302: "Found",
		303: "See Other",
		304: "Not Modified",
		305: "Use Proxy",
		307: "Temporary Redirect",
		308: "Permanent Redirect",
		400: "Bad Request",
		401: "Unauthorized",
		402: "Payment Required",
		403: "Forbidden",
		404: "Not Found",
		405: "Method Not Allowed",
		406: "Not Acceptable",
		407: "Proxy Auth Required",
		408: "Request Timeout",
		409: "Conflict",
		410: "Gone",
		411: "Length Required",
		412: "Precondition Failed",
		413: "Request Entity Too Large",
		414: "Request URI Too Long",
		415: "Unsupported Media Type",
		416: "Request Range Not Satisfiable",
		417: "Expectation Failed",
		418: "Teapot",
		421: "Misdirected Request",
		422: "Unprocessable Entity",
		423: "Locked",
		424: "Failed Dependency",
		425: "Too Early",
		426: "Upgrade Required",
		429: "Too Many Requests",
		431: "Request Header Fields Too Large",
		451: "Unavailable For Legal Reasons",
		500: "Internal Server Error",
		501: "Not Implemented",
		502: "Bad Gateway",
		503: "Service Unavailable",
		504: "Gateway Timeout",
		505: "HTTP Version Not Supported",
		506: "Variant Also Negotiates",
		507: "Insufficient Storage",
		508: "Loop Detected",
		510: "Not Extended",
		511: "Network Authentication Required",
	}

	return statusMessages[status]
}

func (w *ResponseWriterImpl) WriteResponse(c *gin.Context, response WebResponse) {
	response.Status = w.getStatusMessage(response.Code)
	c.JSON(response.Code, response)
}

func (w *ResponseWriterImpl) AbortHttp(c *gin.Context, err error) {
	msg, code := errors.GetErrorMsg(err)

	c.AbortWithStatusJSON(code, WebResponse{
		Status:  w.getStatusMessage(code),
		Code:    code,
		Message: msg,
	})
}

func (w *ResponseWriterImpl) CustomMsgAbortHttp(c *gin.Context, message string, code int) {
	c.AbortWithStatusJSON(code, WebResponse{
		Status:  w.getStatusMessage(code),
		Code:    code,
		Message: message,
	})
}

func MsgTag(f validator.FieldError) string {
	switch f.Tag() {
	case "required":
		return "this field is required"
	case "email":
		return "invalid email format"
	case "oneof":
		switch true {
		case f.Field() == "Privacy":
			return "input must be one of Public, FriendOnly, Private"
		default:
			return "value must be one of the enum"
		}
	case "required_without":
		switch true {
		case f.Field() == "Text":
			return "text is required if file is empty"
		case f.Field() == "File":
			return "file is required if text is empty"
		default:
			return "this field is required if the other one is empty"
		}
	default:
		return f.Error()
	}
}

func (w *ResponseWriterImpl) HttpValidationErr(c *gin.Context, err error) {
	errMap := make(map[string]string)
	for _, val := range err.(validator.ValidationErrors) {
		errMap[val.Field()] = MsgTag(val)
	}

	c.AbortWithStatusJSON(400, WebResponse{
		Status: w.getStatusMessage(400),
		Code:   400,
		Data:   errMap,
	})
}

func (w *ResponseWriterImpl) New404Error(msg string) error {
	return errors.NewError(msg, 404)
}

func (w *ResponseWriterImpl) New403Error(msg string) error {
	return errors.NewError(msg, 403)
}

func (w *ResponseWriterImpl) New401Error(msg string) error {
	return errors.NewError(msg, 401)
}

func (w *ResponseWriterImpl) New409Error(msg string) error {
	return errors.NewError(msg, 409)
}

func (w *ResponseWriterImpl) New400Error(msg string) error {
	return errors.NewError(msg, 400)
}

func (w *ResponseWriterImpl) NewInvalidObjectIdError() error {
	return errors.NewInvalidObjectIdError()
}

func (w *ResponseWriterImpl) New501Error(msg string) error {
	return errors.NewError(msg, 501)
}

func (w *ResponseWriterImpl) Write200Response(c *gin.Context, msg string, data any) {
	w.WriteResponse(c, WebResponse{
		200,
		w.getStatusMessage(200),
		msg,
		data,
	})
}

func (w *ResponseWriterImpl) Write201Response(c *gin.Context, msg string, data any) {
	w.WriteResponse(c, WebResponse{
		201,
		w.getStatusMessage(201),
		msg,
		data,
	})
}
