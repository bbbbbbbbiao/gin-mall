package bootstrape

import (
	"gin-mall/utils"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

/**
 * @author: biao
 * @date: 2025/9/4 10:31
 * @code: 彼方尚有荣光在
 * @description: 自定义验证规则，将其加到gin框架的验证器中
 */

func InitializeValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义验证器
		_ = v.RegisterValidation("mobile", utils.ValidateMobile)

		// 注册自定义json tag 函数
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}
