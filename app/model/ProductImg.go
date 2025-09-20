package model

import "gorm.io/gorm"

/**
 * @author: biao
 * @date: 2025/9/2 11:11
 * @code: 彼方尚有荣光在
 * @description: 商品图片
 */

type ProductImg struct {
	gorm.Model
	ProductId uint `gorm:"not null"`
	ImgPath   string
}
