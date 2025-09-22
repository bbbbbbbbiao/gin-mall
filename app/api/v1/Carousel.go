package v1

import (
	"gin-mall/app/common/response"
	"gin-mall/app/serializer"
	"gin-mall/app/service"
	"github.com/gin-gonic/gin"
)

/**
 * @author: biao
 * @date: 2025/9/21 19:19
 * @code: 彼方尚有荣光在
 * @description: 轮播图Controller
 */

// 获取轮播图
func GetCarousels(c *gin.Context) {

	err, carousels := service.CarouselService.GetCarousels()

	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}

	response.Success(c, serializer.BuildDataList(uint(len(carousels)), serializer.BuildCarouselList(carousels)))
}
