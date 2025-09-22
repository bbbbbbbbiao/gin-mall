package service

import (
	"gin-mall/app/dao"
	"gin-mall/app/model"
	"gin-mall/global"
	"go.uber.org/zap"
)

/**
 * @author: biao
 * @date: 2025/9/21 19:23
 * @code: 彼方尚有荣光在
 * @description: 轮播图服务
 */

type carouselService struct {
}

var CarouselService = new(carouselService)

// 获取轮播图
func (c *carouselService) GetCarousels() (err error, carousels []*model.Carousel) {
	carousels, err = dao.CarouselDao.GetCarousels()

	if err != nil {
		global.App.Log.Error("获取轮播图信息失败", zap.Any("err", err))
		return err, nil
	}

	return err, carousels
}
