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

func (p *productService) CreateProduct(id uint, params request.ProductInfo, files []*multipart.FileHeader) (err error, product *model.Product) {

	boss, err := dao.UserDao.GetUserById(id)
	if err != nil {
		global.App.Log.Error("获取用户信息失败", zap.Any("err", err))
		return err, nil
	}

	// 以第一张图片作为封面
	fileOne, _ := files[0].Open()
	err, path := UploadService.UploadProductToLocalStatic(id, fileOne)
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
	for _, file := range files {

		go func() (err error, product *model.Product) {
			temp, _ := file.Open()

			err, filePath := UploadService.UploadProductToLocalStatic(id, temp)

			if err != nil {
				global.App.Log.Error("上传商品图片至本地失败", zap.Any("err", err))
				return err, nil
			}
			productImg := &model.ProductImg{
				ProductId: product.ID,
				ImgPath:   filePath,
			}

			err = dao.ProductImgDao.CreateProductImg(productImg)
			if err != nil {
				global.App.Log.Error("创建商品图片信息失败", zap.Any("err", err))
				return err, nil
			}
			wg.Done()
			return
		}()
	}

	wg.Wait()
	return
}
