package config

/**
 * @author: biao
 * @date: 2025/9/1 16:08
 * @code: 彼方尚有荣光在
 * @description: 缓存配置结构体
 */

type Redis struct {
	RedisDb     string `mapstructure:"redis_db" json:"redis_db" yaml:"redis_db"`
	RedisAddr   string `mapstructure:"redis_addr" json:"redis_addr" yaml:"redis_addr"`
	RedisPw     string `mapstructure:"redis_pw" json:"redis_pw" yaml:"redis_pw"`
	RedisDbName string `mapstructure:"redis_db_name" json:"redis_db_name" yaml:"redis_db_name"`
}
