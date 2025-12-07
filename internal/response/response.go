package response

import (
	"time"

	"github.com/gin-gonic/gin"
)

const (
	CodeSuccess       = 200
	CodeBadRequest    = 400
	CodeUnauthorized  = 401
	CodeForbidden     = 403
	CodeNotFound      = 404
	CodeInternalError = 500

	CodeParamError = 40001
)

var codeMessageMap = map[init]string{
	CodeSuccess:       "成功",
	CodeBadRequest:    "请求参数错误",
	CodeUnauthorized:  "未授权访问",
	CodeForbidden:     "禁止访问",
	CodeNotFound:      "资源不存在",
	CodeInternalError: "服务器内部错误",

	CodeParamError: "参数错误",
}

type Response struct {
	Code      init        `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp int64       `json:"timestamp"`
	TraceID   string      `json:"trace_id,omitempty"`
}

type Pagination struct {
	Page      int   `json:"page"`
	PageSize  int   `json:"page_size"`
	Total     int64 `json:"total"`
	TotalPage int64 `json:"total_page"`
}

type PageResponse struct {
	List       interface{} `json:"list"`
	Pagination Pagination  `json:"pagination"`
}

type Empty struct{}

func NewResponse() *Response {
	return &Response{
		Timestamp: time.Now().Unix(),
	}
}

func Success(c *gin.Context, data interface{}) {
	resp := NewResponse()
	resp.Code = CodeSuccess
}
