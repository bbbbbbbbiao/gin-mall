package serializer

import (
	"gin-mall/app/model"
	"gin-mall/global"
)

/**
 * @author: biao
 * @date: 2025/9/22 18:21
 * @code: 彼方尚有荣光在
 * @description: 返回的商品信息
 */

type Product struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	CategoryId    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	View          int64  `json:"view"`
	CreateAt      int64  `json:"create_at"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
	BossId        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
	BossAvatar    string `json:"boss_avatar"`
}

func BuildProduct(product *model.Product) *Product {
	return &Product{
		Id:            product.ID,
		Name:          product.Name,
		CategoryId:    product.CategoryId,
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       global.App.Config.Path.Host + global.App.Config.Path.ProductPath + product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		View:          product.GetView(),
		CreateAt:      product.CreatedAt.Unix(),
		Num:           product.Num,
		OnSale:        product.OnSale,
		BossId:        product.BossId,
		BossName:      product.BossName,
		BossAvatar:    global.App.Config.Path.Host + global.App.Config.Path.AvatarPath + product.BossAvatar,
	}
}

func BuildProductList(products []*model.Product) (productList []*Product) {
	for _, product := range products {
		productList = append(productList, BuildProduct(product))
	}
	return
}
