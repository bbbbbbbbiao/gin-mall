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

// 创建商品
func (productDao *productDao) CreateProduct(product *model.Product) (err error) {
	return global.App.DB.Create(product).Error
}

// 计算对应条件下商品个数
func (productDao *productDao) CountProductByCondition(condition map[string]interface{}) (err error, total int64) {
	err = global.App.DB.Model(&model.Product{}).Where(condition).Count(&total).Error
	return
}

// 按条件查询商品
func (productDao *productDao) ProductList(condition map[string]interface{}, page model.BasePage) (productList []*model.Product, err error) {
	err = global.App.DB.Where(condition).Offset((page.PageNum - 1) * page.PageSize).Limit(page.PageSize).Find(&productList).Error
	return
}

// 搜索商品
func (productDao *productDao) ProductSearch(info string, page model.BasePage) (err error, productList []*model.Product, total int64) {
	// 查询个数
	err = global.App.DB.Model(&model.Product{}).
		Where("info like ? or title like ?", "%"+info+"%", "%"+info+"%").
		Count(&total).
		Error
	// 分页获取数据
	err = global.App.DB.Where("info like ? or title like ?", "%"+info+"%", "%"+info+"%").
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).Find(&productList).
		Error
	return
}

// 获取商品信息
func (productDao *productDao) ProductInfoById(id uint) (err error, product *model.Product) {
	err = global.App.DB.Where("id = ?", id).First(&product).Error
	return
}

// 获取商品图片信息
func (productDao *productDao) ProductImgInfoById(id uint) (err error, productImgList []*model.ProductImg) {
	err = global.App.DB.Where("product_id = ?", id).Find(&productImgList).Error
	return
}

// 获取商品分类信息
func (productDao *productDao) Categories() (err error, categoryList []*model.Category) {
	err = global.App.DB.Find(&categoryList).Error
	return
}

// 检查商品是否存在
func (productDao *productDao) ProductIsExist(productId uint, bossId uint) (exist bool, err error) {
	var count int64
	err = global.App.DB.Model(&model.Product{}).Where("id = ? and boss_id = ?", productId, bossId).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
