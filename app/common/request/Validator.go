package request

import "github.com/go-playground/validator/v10"

/**
 * @author: biao
 * @date: 2025/9/4 10:09
 * @code: 彼方尚有荣光在
 * @description: 请求参数的验证
 */

type Validator interface {
	GetMessages() ValidatorMessages
}

type ValidatorMessages map[string]string

func GetErrorMsg(request interface{}, err error) string {
	if _, isValidatorErrors := err.(validator.ValidationErrors); isValidatorErrors {
		_, isValidator := request.(Validator)
		for _, v := range err.(validator.ValidationErrors) {
			if isValidator {
				key := v.Field() + "." + v.Tag()
				if message, ok := request.(Validator).GetMessages()[key]; ok {
					return message
				}
			}
			return v.Error()
		}
	}
	return "Paramter error"
}
