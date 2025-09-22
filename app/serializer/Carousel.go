package serializer

import (
	"gin-mall/app/model"
	"time"
)

/**
 * @author: biao
 * @date: 2025/9/21 19:44
 * @code: 彼方尚有荣光在
 * @description: 轮播图返回格式设置
 */

type Carousel struct {
	Id        uint   `josn:"id"`
	ImgPath   string `josn:"img_path"`
	ProductId uint   `josn:"product_id"`
	CreateAt  int64  `josn:"create_at"`
}

func BuildCarousel(carousel *model.Carousel) Carousel {
	return Carousel{
		Id:        carousel.ID,
		ImgPath:   carousel.ImgPath,
		ProductId: carousel.ProductId,
		CreateAt:  time.Now().Unix(),
	}
}

func BuildCarouselList(carousels []*model.Carousel) (carouselList []Carousel) {
	for _, carousel := range carousels {
		buildCarousel := BuildCarousel(carousel)
		carouselList = append(carouselList, buildCarousel)
	}
	return
}
