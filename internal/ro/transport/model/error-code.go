package model

type ErrorCode struct {
	code int
	desc string
}

var (
	Error200 = build(200, "操作成功")
	Error404 = build(404, "资源未找到")
	Error500 = build(500, "服务异常")
)

func build(code int, desc string) *ErrorCode {
	return &ErrorCode{
		code: code,
		desc: desc,
	}
}

func (e *ErrorCode) IsOk() bool {
	return e.code == 200
}

func (e *ErrorCode) Code() int {
	return e.code
}

func (e *ErrorCode) Desc() string {
	return e.desc
}

func (e *ErrorCode) String() string {
	return e.desc
}
