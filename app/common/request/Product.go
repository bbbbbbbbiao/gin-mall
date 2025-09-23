package request

import "gin-mall/app/model"

/**
 * @author: biao
 * @date: 2025/9/22 15:04
 * @code: 彼方尚有荣光在
 * @description: 商品请求信息
 */

type ProductInfo struct {
	Id            uint   `json:"id" form:"id"`
	Name          string `json:"name" form:"name"`
	CategoryId    uint   `json:"category_id" form:"category_id"`
	Title         string `json:"title" form:"title"`
	Info          string `json:"info" form:"info"`
	ImgePath      string `json:"image_path" form:"image_path"`
	Price         string `json:"price" form:"price"`
	DiscountPrice string `json:"discount_price" form:"discount_price"`
	OnSale        string `json:"on_sale" form:"on_sale"`
	Num           int    `json:"num" form:"num"`
	model.BasePage
}
