package v1

import (
	"gin-mall/app/common/request"
	"gin-mall/app/common/response"
	"gin-mall/app/service"
	"gin-mall/global"
	"gin-mall/serializer"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/**
 * @author: biao
 * @date: 2025/9/2 20:50
 * @code: 彼方尚有荣光在
 * @description: 用户Controller
 */

// 用户注册
func UserRegister(c *gin.Context) {
	var userRegister request.UserRegister

	// 用户信息绑定
	if err := c.ShouldBind(&userRegister); err != nil {
		errMsg := request.GetErrorMsg(userRegister, err)
		global.App.Log.Error("Register-用户请求信息失败：", zap.String("err", errMsg))
		response.ValidateFail(c, errMsg)
		return
	}

	// 用户信息注册
	if err, user := service.UserService.UserRegister(userRegister); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}

// 用户登录
func UserLogin(c *gin.Context) {
	var userLogin request.UserLogin

	// 用户信息绑定
	if err := c.ShouldBind(&userLogin); err != nil {
		errMsg := request.GetErrorMsg(userLogin, err)
		global.App.Log.Error("Login-用户请求信息失败：", zap.String("err", errMsg))
		response.ValidateFail(c, errMsg)
		return
	}

	// 用户登录
	err, user := service.UserService.UserLogin(userLogin)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}

	// 签发 Token
	tokenData, err, _ := service.JwtService.CreateToken(user.ID, user.UserName, 0, service.AppGuardName)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}

	response.Success(c, serializer.BuildUserAndToken(user, tokenData))
}

// 用户修改信息
func UserUpdateInfo(c *gin.Context) {
	userUpdateInfo := request.UserUpdateInfo{}

	if err := c.ShouldBind(&userUpdateInfo); err != nil {
		response.BusinessFail(c, err.Error())
	}
	v, _ := c.Get("id")

	err, user := service.UserService.UserUpdateInfo(userUpdateInfo, v.(uint))

	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, serializer.BuildUser(user))
}

// 上传头像至本地
func UploadAvatarToLocal(c *gin.Context) {
	file, _, _ := c.Request.FormFile("file")

	v, _ := c.Get("id")
	err, path := service.UserService.UploadAvatarToLocal(v.(uint), file)

	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, path)
}

// 发送邮件
func SendEmail(c *gin.Context) {
	var sendEmail request.SendEmail

	if err := c.ShouldBind(&sendEmail); err != nil {
		global.App.Log.Error("Login-用户请求信息失败：", zap.Any("err", err))
		response.BusinessFail(c, err.Error())
	}

	v, _ := c.Get("id")
	err := service.UserService.SendEmail(v.(uint), sendEmail)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}

	response.Success(c, sendEmail.Email+": 发送邮件成功")
}

// 验证邮件
func ValidEmail(c *gin.Context) {
	err, user := service.UserService.ValidEmail(c)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, serializer.BuildUser(user))
}
