package service

import (
	"gin-mall/app/common/request"
	"gin-mall/app/dao"
	"gin-mall/app/model"
	"gin-mall/global"
	"go.uber.org/zap"
	"mime/multipart"
	"sync"
)

/**
 * @author: biao
 * @date: 2025/9/22 15:16
 * @code: 彼方尚有荣光在
 * @description: produce 处理服务
 */

type productService struct {
}

var ProductService = new(productService)

// 创建商品
func (p *productService) CreateProduct(id uint, params request.ProductInfo, files []*multipart.FileHeader) (err error, product *model.Product) {

	boss, err := dao.UserDao.GetUserById(id)
	if err != nil {
		global.App.Log.Error("获取用户信息失败", zap.Any("err", err))
		return err, nil
	}

	// 以第一张图片作为封面
	fileOne, _ := files[0].Open()
	err, path := UploadService.UploadProductToLocalStatic(id, fileOne, 0)
	if err != nil {
		global.App.Log.Error("上传商品图片至本地失败", zap.Any("err", err))
		return err, nil
	}

	product = &model.Product{
		Name:          params.Name,
		CategoryId:    params.CategoryId,
		Title:         params.Title,
		Info:          params.Info,
		ImgPath:       path,
		DiscountPrice: params.DiscountPrice,
		OnSale:        true,
		Num:           params.Num,
		BossId:        id,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}

	err = dao.ProductDao.CreateProduct(product)
	if err != nil {
		global.App.Log.Error("创建商品信息失败", zap.Any("err", err))
		return err, nil
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	index := 0
	for _, file := range files {
		index++
		go func(index int) (err error) {
			temp, _ := file.Open()

			err, filePath := UploadService.UploadProductToLocalStatic(id, temp, index)

			if err != nil {
				global.App.Log.Error("上传商品图片至本地失败", zap.Any("err", err))
				return err
			}
			productImg := &model.ProductImg{
				ProductId: product.ID,
				ImgPath:   filePath,
			}

			err = dao.ProductImgDao.CreateProductImg(productImg)
			if err != nil {
				global.App.Log.Error("创建商品图片信息失败", zap.Any("err", err))
				return err
			}
			wg.Done()
			return
		}(index)

	}

	wg.Wait()
	return
}

// 商品列表
func (p *productService) ProductList(params request.ProductInfo) (err error, productList []*model.Product, total int64) {
	condition := make(map[string]interface{})

	if params.PageSize == 0 {
		params.PageSize = 15
	}
	if params.CategoryId != 0 {
		condition["category_id"] = params.CategoryId
	}

	err, total = dao.ProductDao.CountProductByCondition(condition)
	if err != nil {
		global.App.Log.Error("查询商品数量失败", zap.Any("err", err))
		return err, nil, 0
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productList, _ = dao.ProductDao.ProductList(condition, params.BasePage)
		wg.Done()
	}()
	wg.Wait()
	return
}

// 商品搜索
func (p *productService) ProductSearch(params request.ProductInfo) (err error, productList []*model.Product, total int64) {
	if params.PageSize == 0 {
		params.PageSize = 15
	}

	err, productList, total = dao.ProductDao.ProductSearch(params.Info, params.BasePage)
	if err != nil {
		global.App.Log.Error("商品信息搜索失败", zap.Any("err", err))
		return err, nil, 0
	}

	return
}

// 获取商品信息
func (p *productService) ProductInfoById(id uint) (err error, product *model.Product) {
	err, product = dao.ProductDao.ProductInfoById(id)
	if err != nil {
		global.App.Log.Error("获取商品信息失败，id：", zap.Uint("id", id))
		return
	}
	return
}

// 获取商品图片信息
func (p *productService) ProductImgById(id uint) (err error, productImgList []*model.ProductImg) {
	err, productImgList = dao.ProductDao.ProductImgInfoById(id)
	if err != nil {
		global.App.Log.Error("获取商品图片信息失败，id：", zap.Uint("id", id))
		return
	}
	return
}

func (p *productService) Categories() (err error, categoryList []*model.Category) {
	err, categoryList = dao.ProductDao.Categories()
	if err != nil {
		global.App.Log.Error("获取商品分类信息失败", zap.Any("err", err))
		return
	}
	return
}
