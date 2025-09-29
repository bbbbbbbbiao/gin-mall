package v1

import (
	"gin-mall/app/common/request"
	"gin-mall/app/common/response"
	"gin-mall/app/serializer"
	"gin-mall/app/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

/**
 * @author: biao
 * @date: 2025/9/28 下午5:59
 * @code: 彼方尚有荣光在
 * @description: 收藏夹Controller
 */

// FavoriteAdd 添加收藏夹
func FavoriteAdd(c *gin.Context) {
	var favoriteAdd request.Favorite

	if err := c.ShouldBind(&favoriteAdd); err != nil {
		response.BusinessFail(c, err.Error())
		return
	}

	id, _ := c.Get("id")
	err := service.FavoriteService.FavoriteAdd(id.(uint), favoriteAdd)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}

	response.Success(c, "添加收藏成功")
}

// FavoriteList 获取收藏夹列表
func FavoriteList(c *gin.Context) {
	id, _ := c.Get("id")
	FavoriteList, err := service.FavoriteService.FavoriteList(id.(uint))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, serializer.BuildDataList(uint(len(FavoriteList)), serializer.BuildFavoriteList(FavoriteList)))
}

// FavoriteDelete 删除收藏夹
func FavoriteDelete(c *gin.Context) {

	id, _ := c.Get("id")
	productId, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	err := service.FavoriteService.FavoriteDelete(id.(uint), uint(productId))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}

	response.Success(c, "删除收藏成功")
}
