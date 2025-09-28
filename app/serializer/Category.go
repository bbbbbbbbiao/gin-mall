package serializer

import "gin-mall/app/model"

/**
 * @author: biao
 * @date: 2025/9/27 下午12:10
 * @code: 彼方尚有荣光在
 * @description: 商品分类展示序列化器
 */

type Category struct {
	Id           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	CreateAt     int64  `json:"create_at"`
}

func BuildCategory(category *model.Category) *Category {
	return &Category{
		Id:           category.ID,
		CategoryName: category.CatrgoryName,
		CreateAt:     category.Model.CreatedAt.Unix(),
	}
}

func BuildCategoryList(categories []*model.Category) (categoryList []*Category) {
	for _, category := range categories {
		categoryList = append(categoryList, BuildCategory(category))
	}
	return
}
