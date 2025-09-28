package serializer

import "gin-mall/app/model"

/**
 * @author: biao
 * @date: 2025/9/26 10:02
 * @code: 彼方尚有荣光在
 * @description: 返回的商品图片信息
 */

type ProductImg struct {
	ProductId uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
}

func BuildProductImg(productImg *model.ProductImg) *ProductImg {
	return &ProductImg{
		ProductId: productImg.ProductId,
		ImgPath:   productImg.ImgPath,
	}
}

func BuildProductImgList(productImgs []*model.ProductImg) (productImgList []*ProductImg) {
	for _, productImg := range productImgs {
		productImgList = append(productImgList, BuildProductImg(productImg))
	}
	return
}
