package model

import (
	"context"
	"gin-mall/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
)

/**
 * @author: biao
 * @date: 2025/9/2 11:01
 * @code: 彼方尚有荣光在
 * @description: 商品信息
 */

type Product struct {
	gorm.Model
	Name          string
	CategoryId    uint
	Title         string // 商品标题
	Info          string // 商品详情信息
	ImgPath       string // 图片地址
	Prince        string //价格
	DiscountPrice string // 折扣后的价格
	OnSale        bool   // 是否在售
	Num           int    // 数量
	BossId        uint   // 商家id
	BossName      string // 商家名称
	BossAvatar    string // 商家头像
}

func (p *Product) GetProductKey() string {
	return "product_key_" + strconv.Itoa(int(p.ID))
}

func (p *Product) GetView() int64 {
	res, err := global.App.Redis.Get(context.Background(), p.GetProductKey()).Result()
	resInt, err := strconv.ParseInt(res, 10, 64)
	if err != nil {
		global.App.Log.Error("Redis-获取商品的浏览量失败", zap.Any("err", err))
		return -1
	}

	return resInt
}

func (p *Product) AddView() {
	_, err := global.App.Redis.Incr(context.Background(), p.GetProductKey()).Result()

	if err != nil {
		global.App.Log.Error("Redis-商品浏览量增加失败", zap.Any("err", err))
	}
}
