package model

import "gorm.io/gorm"

/**
 * @author: biao
 * @date: 2025/9/2 9:47
 * @code: 彼方尚有荣光在
 * @description: 管理员信息
 */

type Admin struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	Avatar         string
}
