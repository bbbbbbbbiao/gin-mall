package dao

import (
	"gin-mall/app/model"
	"gin-mall/global"
)

/**
 * @author: biao
 * @date: 2025/9/22 16:36
 * @code: 彼方尚有荣光在
 * @description: 操作ProductImg数据库
 */

type productImgDao struct {
}

var ProductImgDao = new(productImgDao)

func (productImgDao *productImgDao) CreateProductImg(productImg *model.ProductImg) (err error) {
	return global.App.DB.Create(productImg).Error
}
