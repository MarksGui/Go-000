package main

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(code int64, msg string) *Error {
	return &Error{Code: code, Message: msg}
}

var (
	ErrOk       = NewError(0, "成功")
	ErrServer   = NewError(1000, "服务错误")
	ErrNotFound = NewError(2001, "查无记录")
)
