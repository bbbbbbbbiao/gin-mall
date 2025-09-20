package utils

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

/**
 * @author: biao
 * @date: 2025/9/4 10:28
 * @code: 彼方尚有荣光在
 * @description: 验证器工具类，自定义的验证规则都写在这
 */

// 验证手机号码
func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	ok, _ := regexp.MatchString(`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`, mobile)
	if !ok {
		return false
	}
	return true
}
