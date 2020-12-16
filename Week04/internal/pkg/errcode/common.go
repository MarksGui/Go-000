package errcode

var (
	Success          = NewError(0, "成功")
	ErrServer        = NewError(1001, "服务内部错误")
	ErrInvalidParams = NewError(1002, "入参错误")
	ErrNotFound      = NewError(1003, "找不到")
	ErrDatabase      = NewError(1004, "数据库错误")
)
