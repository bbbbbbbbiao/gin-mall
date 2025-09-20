package model

import (
	"gorm.io/gorm"
)

/**
 * @author: biao
 * @date: 2025/9/2 11:01
 * @code: 彼方尚有荣光在
 * @description: 商品信息
 */

type Product struct {
	gorm.Model
	Name          string
	Catrgory      uint
	Title         string // 商品标题
	Info          string // 商品详情信息
	ImgPath       string // 图片地址
	DiscountPrice string // 折扣后的价格
	OnSale        bool   // 是否在售
	Num           int    // 数量
	BoosId        uint   // 商家id
	BossName      string // 商家名称
	BossAvatar    string // 商家头像
}
