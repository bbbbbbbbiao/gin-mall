package model

import "gorm.io/gorm"

/**
 * @author: biao
 * @date: 2025/9/2 10:42
 * @code: 彼方尚有荣光在
 * @description: 轮播图信息
 */

type Carousel struct {
	gorm.Model
	ImgPath   string
	ProductId uint `gorm:"not null"`
}
