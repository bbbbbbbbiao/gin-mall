package model

import "gorm.io/gorm"

/**
 * @author: biao
 * @date: 2025/9/2 10:50
 * @code: 彼方尚有荣光在
 * @description: 商品种类信息
 */

type Category struct {
	gorm.Model
	CatrgoryName string
}
