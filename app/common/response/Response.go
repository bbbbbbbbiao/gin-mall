package response

import (
	"gin-mall/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @author: biao
 * @date: 2025/9/2 21:20
 * @code: 彼方尚有荣光在
 * @description: 封装响应
 */

type Response struct {
	ErrorCode int         `json:"error_code"`
	Data      interface{} `json:"data"`
	Message   string      `json:"message"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		ErrorCode: 0,
		Data:      data,
		Message:   "ok",
	})
}

func Fail(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Response{
		ErrorCode: code,
		Data:      nil,
		Message:   message,
	})
}

func FailByError(c *gin.Context, err global.CustomError) {
	Fail(c, err.ErrorCode, err.ErrorMsg)
}

// ValidateFail 请求参数验证失败
func ValidateFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.ValidateError.ErrorCode, msg)
}

// BusinessFail 业务逻辑失败
func BusinessFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.BusinessError.ErrorCode, msg)
}

// TokenFail Token鉴权失败
func TokenFail(c *gin.Context) {
	FailByError(c, global.Errors.TokenError)
}

// EmailTokenFail EmailToken鉴权失败
func EmailTokenFail(c *gin.Context) {
	FailByError(c, global.Errors.EmailTokenError)
}
