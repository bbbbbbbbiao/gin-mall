package dao

import (
	"gin-mall/app/model"
	"gin-mall/global"
)

/**
 * @author: biao
 * @date: 2025/9/28 下午6:15
 * @code: 彼方尚有荣光在
 * @description: 收藏商品的数据库操作
 */

type favoriteDao struct {
}

var FavoriteDao = new(favoriteDao)

// 检查收藏夹中商品是否存在
func (favoriteDao *favoriteDao) ProductIsExist(userId uint, productId uint) (bool, error) {
	var count int64
	err := global.App.DB.Model(&model.Favorite{}).Where("user_id = ? and product_id = ?", userId, productId).Count(&count).Error

	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// 添加收藏夹商品
func (favoriteDao *favoriteDao) FavoriteAdd(favorite *model.Favorite) error {
	return global.App.DB.Create(&favorite).Error
}

func (favoriteDao *favoriteDao) FavoriteList(id uint) (favoriteList []*model.Favorite, err error) {
	err = global.App.DB.Where("user_id = ?", id).Find(&favoriteList).Error
	return favoriteList, err
}

// FavoriteDelete 删除收藏夹商品
func (favoriteDao *favoriteDao) FavoriteDelete(id uint, productId uint) error {
	return global.App.DB.Model(&model.Favorite{}).Where("user_id = ? and product_id = ?", id, productId).Delete(&model.Favorite{}).Error
}
