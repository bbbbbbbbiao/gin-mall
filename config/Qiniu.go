package config

/**
 * @author: biao
 * @date: 2025/9/1 16:09
 * @code: 彼方尚有荣光在
 * @description: 七牛云配置结构体
 */

type Qiniu struct {
	AccessKey   string `mapstructure:"access_key" json:"access_key" yaml:"access_key"`
	SecretKey   string `mapstructure:"secret_key" json:"secret_key" yaml:"secret_key"`
	Bucket      string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	QiniuServer string `mapstructure:"qiniu_server" json:"qiniu_server" yaml:"qiniu_server"`
}
