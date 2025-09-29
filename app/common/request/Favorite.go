package request

import "gin-mall/app/model"

/**
 * @author: biao
 * @date: 2025/9/28 下午5:55
 * @code: 彼方尚有荣光在
 * @description: 收藏夹请求参数
 */

type Favorite struct {
	UserId    uint `json:"user_id" form:"user_id"`
	ProductId uint `json:"product_id" form:"product_id"`
	BossId    uint `json:"boss_id" form:"boss_id"`
	model.BasePage
}
