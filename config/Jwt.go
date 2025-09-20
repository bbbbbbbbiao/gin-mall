package config

/**
 * @author: biao
 * @date: 2025/9/8 10:50
 * @code: 彼方尚有荣光在
 * @description: JWT配置结构体
 */

type Jwt struct {
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"`
	JwtTtl int64  `mapstructure:"jwt_ttl" json:"jwt_ttl" yaml:"jwt_ttl"`
}
