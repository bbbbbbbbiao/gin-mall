package model

import "gorm.io/gorm"

/**
 * @author: biao
 * @date: 2025/9/2 10:54
 * @code: 彼方尚有荣光在
 * @description: 收藏夹信息
 */

type Favorite struct {
	gorm.Model
	User      User
	UserId    uint `gorm:"not null"`
	Product   Product
	ProductId uint `gorm:"not null"`
	Boss      User
	BossId    uint `gorm:"not null"`
}
