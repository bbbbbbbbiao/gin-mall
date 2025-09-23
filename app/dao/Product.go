package dao

import (
	"gin-mall/app/model"
	"gin-mall/global"
)

/**
 * @author: biao
 * @date: 2025/9/22 16:29
 * @code: 彼方尚有荣光在
 * @description: 操作Product数据库
 */

type productDao struct {
}

var ProductDao = new(productDao)

func (productDao *productDao) CreateProduct(product *model.Product) (err error) {
	return global.App.DB.Create(product).Error
}
