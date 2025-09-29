package serializer

import "gin-mall/app/model"

/**
 * @author: biao
 * @date: 2025/9/28 下午7:04
 * @code: 彼方尚有荣光在
 * @description: 收藏夹序列化器
 */

type Favorite struct {
	UserId        uint   `json:"user_id"`
	ProductId     uint   `json:"product_id"`
	Name          string `json:"name"`
	CategoryId    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	BossId        uint   `json:"boss_id"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
}

func BuildFavorite(favorite *model.Favorite) *Favorite {
	return &Favorite{
		UserId:        favorite.UserId,
		ProductId:     favorite.ProductId,
		Name:          favorite.Product.Name,
		CategoryId:    favorite.Product.CategoryId,
		Title:         favorite.Product.Title,
		Info:          favorite.Product.Info,
		ImgPath:       favorite.Product.ImgPath,
		Price:         favorite.Product.Price,
		DiscountPrice: favorite.Product.DiscountPrice,
		BossId:        favorite.BossId,
		Num:           favorite.Product.Num,
		OnSale:        favorite.Product.OnSale,
	}
}

func BuildFavoriteList(favorites []*model.Favorite) (favoriteList []*Favorite) {
	for _, favorite := range favorites {
		favoriteList = append(favoriteList, BuildFavorite(favorite))
	}
	return
}
