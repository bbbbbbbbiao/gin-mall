package bootstrape

import (
	"context"
	"gin-mall/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

/**
 * @author: biao
 * @date: 2025/9/23 11:08
 * @code: 彼方尚有荣光在
 * @description: 初始化Redis
 */

func InitializeRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     global.App.Config.Redis.RedisAddr,
		Password: global.App.Config.Redis.RedisPw,
		DB:       global.App.Config.Redis.RedisDbName,
	})

	_, err := client.Ping(context.Background()).Result()

	if err != nil {
		global.App.Log.Error("Redis 连接失败", zap.Any("err", err))
		return nil
	}

	return client
}
