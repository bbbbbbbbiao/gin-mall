package dao

import (
	model2 "gin-mall/app/model"
	"gin-mall/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
)

/**
 * @author: biao
 * @date: 2025/9/1 21:25
 * @code: 彼方尚有荣光在
 * @description: 数据库迁移
 */

func Migration(db *gorm.DB) {
	err := db.AutoMigrate(
		&model2.User{},
		&model2.Address{},
		&model2.Admin{},
		&model2.Catrgory{},
		&model2.Carousel{},
		&model2.Cart{},
		&model2.Notice{},
		&model2.Product{},
		&model2.ProductImg{},
		&model2.Order{},
		&model2.Favorite{},
	)

	if err != nil {
		global.App.Log.Error("migrate table failed, err : ", zap.Any("err", err))
		os.Exit(0)
	}
}
