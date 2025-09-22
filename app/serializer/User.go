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

type Money struct {
	UserId    uint   `json:"user_id" form:"user_id"`
	UserName  string `form:"user_name" json:"user_name"`
	UserMoney string `form:"user_money" json:"user_money"`
}

func BuildToken(user *model.User, tokenOutPut *service.TokenOutPut) map[string]interface{} {
	user.Avatar = global.App.Config.Path.Host + global.App.Config.Path.AvatarPath + user.Avatar
	return map[string]interface{}{
		"user":      user,
		"tokenData": tokenOutPut,
	}
}

func BuildAvatar(user *model.User) *model.User {
	user.Avatar = global.App.Config.Path.Host + global.App.Config.Path.AvatarPath + user.Avatar
	return user
}

func BuildMoney(user *model.User, money string) Money {
	return Money{
		UserId:    user.ID,
		UserName:  user.UserName,
		UserMoney: money,
	}
}
