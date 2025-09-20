package model

import "gorm.io/gorm"

/**
 * @author: biao
 * @date: 2025/9/2 10:57
 * @code: 彼方尚有荣光在
 * @description: 用户/商家信息
 */

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"` //账号
	Email          string // 邮箱
	PasswordDigest string // 密码
	NickName       string // 昵称
	Status         string // 状态：（卖违规产品）判断是否被封禁
	Avatar         string // 头像
	Money          string
}

const (
	PasswordCost        = 12       // 密码加密难度
	Active       string = "active" //激活用户
)
