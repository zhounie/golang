package e

import "fmt"

type AppError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Code:%d, Msg:%s", e.Code, e.Msg)
}

func New(code int, msg string) *AppError {
	return &AppError{Code: code, Msg: msg}
}

var (
	CodeSuccess       = New(200, "成功")
	CodeBadRequest    = New(400, "请求参数错误")
	CodeUnauthorized  = New(401, "未授权访问")
	CodeForbidden     = New(403, "禁止访问")
	CodeNotFound      = New(404, "资源不存在")
	CodeInternalError = New(500, "服务器内部错误")
)
