package config

/**
 * @author: biao
 * @date: 2025/9/1 15:21
 * @code: 彼方尚有荣光在
 * @description: 读取配置文件
 */

type Configuration struct {
	Service Service `mapstructure:"service" json:"service" yaml:"service"`
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Qiniu   Qiniu   `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	Email   Email   `mapstructure:"email" json:"email" yaml:"email"`
	Path    Path    `mapstructure:"path" json:"path" yaml:"path"`
	Log     Log     `mapstructure:"log" json:"log" yaml:"log"`
	Jwt     Jwt     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
