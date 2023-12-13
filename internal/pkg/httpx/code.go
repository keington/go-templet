package httpx

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/7 23:22
 * @file: code.go
 * @description: 状态码
 */
// ------------------错误码定义-------------------------------
// 错误码为 5 位数:
// ----------------------------------------------------------
//
//	第1位			2、3位			4、5位			e.g.
//
// ----------------------------------------------------------
//
//	服务级错误码		模块级错误码		具体错误码
//
// ----------------------------------------------------------
//
//	系统级错误		用户模块			手机号不合法		10001
//	业务模块错误		作业模块			参数有误			10002
//	...				...				...
var (
	Failed = failed(500, "failed")

	Unauthorized           = failed(4001, "unauthorized")
	AuthorizationIncorrect = failed(4002, "The auth format in the request header is incorrect")
	AuthorizationEmpty     = failed(4003, "Authorization is empty")
	TokenInvalid           = failed(4004, "Token is invalid")
	TokenEmpty             = failed(4005, "Token is empty")

	ErrUserPhone = failed(10001, "用户手机号不合法")
	ErrSignParam = failed(10002, "签名参数有误")

	InternalError = failed(5000, "Internal Error")
)

var (
	Success = success(200, "success")
)

// failed 构造函数
func failed(code int, msg string) *Response {
	return &Response{
		Code:   code,
		Msg:    msg,
		Detail: nil,
	}
}

// success 构造函数
func success(code int, msg string) *Response {
	return &Response{
		Code:   code,
		Msg:    msg,
		Detail: nil,
	}
}
