package global

import (
	"gin-mall/config"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

/**
 * @author: biao
 * @date: 2025/9/1 16:43
 * @code: 彼方尚有荣光在
 * @description: 定义全局变量
 */

type Application struct {
	ConfigViper *viper.Viper
	Config      *config.Configuration
	Log         *zap.Logger
	DB          *gorm.DB
	Redis       *redis.Client
}

var App = new(Application)
