package routes

import (
	"gin-mall/app/api/v1"
	"gin-mall/app/middleware"
	"gin-mall/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @author: biao
 * @date: 2025/9/2 20:19
 * @code: 彼方尚有荣光在
 * @description: api路由组
 */

func ApiV1Router(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	// 用户注册
	routerGroup.POST("/user/register", v1.UserRegister)
	// 用户登录
	routerGroup.POST("/user/login", v1.UserLogin)

	// 中间件鉴权
	userGroup := routerGroup.Group("").Use(middleware.ParseJWtAuth(service.AppGuardName))
	{
		// 修改用户信息
		userGroup.POST("/user/update", v1.UserUpdateInfo)
		// 上传图片
		userGroup.POST("/user/uploadAvatar", v1.UploadAvatarToLocal)
		// 发送邮件
		userGroup.POST("/user/sendEmail", v1.SendEmail)
		//显示金额
		userGroup.POST("/user/showMoney", v1.ShowMoney)
	}

	// 验证邮箱
	routerGroup.POST("/user/validEmail", middleware.ParseEmailJWTAuth(service.AppGuardName), v1.ValidEmail)

	// 获取轮播图
	routerGroup.GET("/carousels", v1.GetCarousels)
}
