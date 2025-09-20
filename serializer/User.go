package serializer

import (
	"gin-mall/app/model"
	"gin-mall/app/service"
	"gin-mall/global"
)

/**
 * @author: biao
 * @date: 2025/9/9 11:09
 * @code: 彼方尚有荣光在
 * @description: 给前端所要展示的User数据
 */

func BuildUserAndToken(user *model.User, tokenOutPut *service.TokenOutPut) map[string]interface{} {
	user.Avatar = global.App.Config.Path.Host + global.App.Config.Path.AvatarPath + user.Avatar
	return map[string]interface{}{
		"user":      user,
		"tokenData": tokenOutPut,
	}
}

func BuildUser(user *model.User) *model.User {
	user.Avatar = global.App.Config.Path.Host + global.App.Config.Path.AvatarPath + user.Avatar
	return user
}
