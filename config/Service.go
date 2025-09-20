package config

/**
 * @author: biao
 * @date: 2025/9/1 16:08
 * @code: 彼方尚有荣光在
 * @description: 服务配置
 */

type Service struct {
	AppMode  string `mapstructure:"app_mode" json:"app_mode" yaml:"app_mode"`
	HttpHost string `mapstructure:"http_host" json:"http_host" yaml:"http_host"`
	HttpPort string `mapstructure:"http_port" json:"http_port" yaml:"http_port"`
}
