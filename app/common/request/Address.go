package request

/**
 * @author: biao
 * @date: 2025/9/29 上午11:15
 * @code: 彼方尚有荣光在
 * @description: 地址请求参数
 */

type Address struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Phone   string `json:"phone" form:"phone" binding:"required"`
	Address string `json:"address" form:"address" binding:"required"`
}

func (a *Address) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required":    "姓名不能为空",
		"phone.required":   "手机号不能为空",
		"address.required": "地址不能为空",
	}
}
