package model

import "gorm.io/gorm"

/**
 * @author: biao
 * @date: 2025/9/2 10:44
 * @code: 彼方尚有荣光在
 * @description: 购物车信息
 */

type Cart struct {
	gorm.Model
	UsertId   uint `gorm:"not null"`      // 用户id
	ProductId uint `gorm:"not null"`      // 商品id
	BossId    uint `gorm:"not null"`      // 商家id
	Num       int  `gorm:"not null"`      // 商品数量
	MaxNum    int  `gorm:"not null"`      // 该商品的限额数
	Check     bool `gorm:"default:false"` // 是否支付
}
