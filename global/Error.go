package global

/**
 * @author: biao
 * @date: 2025/9/2 21:17
 * @code: 彼方尚有荣光在
 * @description: 自定义错误码
 */

type CustomError struct {
	ErrorCode int
	ErrorMsg  string
}

type CustomErrors struct {
	BusinessError CustomError
	ValidateError CustomError

	UserNameExistError CustomError
	TokenError         CustomError
	EmailTokenError    CustomError
}

var Errors = CustomErrors{
	BusinessError:      CustomError{40000, "业务错误"},
	ValidateError:      CustomError{42200, "请求参数错误"},
	UserNameExistError: CustomError{30001, "用户名已存在错误"},
	TokenError:         CustomError{50100, "登录授权失败"},
	EmailTokenError:    CustomError{50200, "邮箱授权失败"},
}
