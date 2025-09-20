package model

import "gorm.io/gorm"

/**
 * @author: biao
 * @date: 2025/9/2 11:14
 * @code: 彼方尚有荣光在
 * @description: 订单信息
 */

type Order struct {
	gorm.Model
	UserId    uint  `gorm:"not null"`
	ProductId uint  `gorm:"not null"`
	BossId    uint  `gorm:"not null"`
	AddressId uint  `gorm:"not null"`
	Num       int   // 商品数量
	OrderNum  int64 // 订单数量
	Type      uint  // 订单类型 1 未支付 2 已支付
	Money     float64
}
