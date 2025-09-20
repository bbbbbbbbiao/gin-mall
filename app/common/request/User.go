package request

/**
 * @author: biao
 * @date: 2025/9/2 21:29
 * @code: 彼方尚有荣光在
 * @description: User信息
 */

// 注册
type UserRegister struct {
	NickName string `json:"nick_name" form:"nick_name" binding:"required"`
	UserName string `json:"user_name" form:"user_name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Key      string `json:"key" form:"key"`
}

func (user UserRegister) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"nick_name.required": "用户昵称不能为空",
		"user_name.required": "用户名不能为空",
		"password.required":  "用户密码不能为空",
	}
}

// 登录
type UserLogin struct {
	UserName string `json:"user_name" form:"user_name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (user UserLogin) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"user_name.required": "用户名不能为空",
		"password.required":  "用户密码不能为空",
	}
}

// 修改信息
type UserUpdateInfo struct {
	NickName string `json:"nick_name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// 发送邮件
type SendEmail struct {
	Email         string `json:"email" form:"email"`
	Password      string `josn:"password" form:"password"`
	OperationType uint   `json:"operation_type" form:"operation_type"`
	// 1 绑定邮箱， 2 解绑邮箱， 3 修改密码
}
