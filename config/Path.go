package config

/**
 * @author: biao
 * @date: 2025/9/1 16:09
 * @code: 彼方尚有荣光在
 * @description: 文件保存路径配置结构体
 */

type Path struct {
	Host        string `mapstructure:"host" json:"host" yaml:"host"`
	ProductPath string `mapstructure:"product_path" json:"product_path" yaml:"product_path"`
	AvatarPath  string `mapstructure:"avatar_path" json:"avatar_path" yaml:"avatar_path"`
}
