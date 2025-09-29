package service

import (
	"gin-mall/app/common/request"
	"gin-mall/app/dao"
	"gin-mall/app/model"
	"gin-mall/global"
	"go.uber.org/zap"
)

/**
 * @author: biao
 * @date: 2025/9/28 下午6:06
 * @code: 彼方尚有荣光在
 * @description: 收藏夹Service
 */

type favoriteService struct {
}

var FavoriteService = new(favoriteService)

// FavoriteAdd 添加收藏夹
func (f *favoriteService) FavoriteAdd(id uint, params request.Favorite) (err error) {

	// 检查商品是否存在
	exist, err := dao.FavoriteDao.ProductIsExist(id, params.ProductId)
	if err != nil {
		global.App.Log.Error("检查商品是否存在失败", zap.Error(err))
		return err
	}
	if exist {
		global.App.Log.Info("商品已存在于收藏夹", zap.Uint("userId", id), zap.Uint("productId", params.ProductId))
		return nil
	}

	// 添加商品到收藏夹
	favorite := &model.Favorite{
		UserId:    id,
		ProductId: params.ProductId,
		BossId:    params.BossId,
	}

	err = dao.FavoriteDao.FavoriteAdd(favorite)
	if err != nil {
		global.App.Log.Error("添加收藏夹失败", zap.Error(err))
		return err
	}

	return nil
}

// FavoriteList 获取收藏夹列表
func (f *favoriteService) FavoriteList(id uint) (favoriteList []*model.Favorite, err error) {
	favoriteList, err = dao.FavoriteDao.FavoriteList(id)
	if err != nil {
		global.App.Log.Error("获取收藏夹列表失败", zap.Error(err))
		return nil, err
	}
	// 获取商品信息
	for i, item := range favoriteList {
		err, product := dao.ProductDao.ProductInfoById(item.ProductId)
		if err != nil {
			global.App.Log.Error("获取商品信息失败", zap.Error(err))
			return nil, err
		}
		favoriteList[i].Product = *product
	}
	return favoriteList, nil
}

// FavoriteDelete 删除收藏夹
func (f *favoriteService) FavoriteDelete(id uint, productId uint) error {
	err := dao.FavoriteDao.FavoriteDelete(id, productId)
	if err != nil {
		global.App.Log.Error("删除收藏夹失败", zap.Error(err))
		return err
	}
	return nil
}

//
