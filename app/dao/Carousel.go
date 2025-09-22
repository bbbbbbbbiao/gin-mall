package dao

import (
	"gin-mall/app/model"
	"gin-mall/global"
)

/**
 * @author: biao
 * @date: 2025/9/21 19:26
 * @code: 彼方尚有荣光在
 * @description: 操作Carousel数据库
 */

type carouselDao struct {
}

var CarouselDao = new(carouselDao)

// 获取轮播图
func (c *carouselDao) GetCarousels() (carousels []*model.Carousel, err error) {
	err = global.App.DB.Model(&model.Carousel{}).Find(&carousels).Error
	return
}
