package model

import "gorm.io/gorm"

/**
 * @author: biao
 * @date: 2025/9/2 9:43
 * @code: 彼方尚有荣光在
 * @description: 用户/商家地址信息
 */

type Address struct {
	gorm.Model
	UserID  uint   `gorm:"not null"`
	Name    string `gorm:"type:varchar(20) not null"` // 名字
	Phone   string `gorm:"type:varchar(11) not null"` // 电话
	Address string `gorm:"type:varchar(50) not null"` // 地址
}
