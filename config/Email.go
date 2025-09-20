package config

/**
 * @author: biao
 * @date: 2025/9/1 16:09
 * @code: 彼方尚有荣光在
 * @description: 邮箱配置结构体
 */

type Email struct {
	ValidEmail string `mapstructure:"valid_email" json:"valid_email" yaml:"valid_email"`
	SmtpHost   string `mapstructure:"smtp_host" json:"smtp_host" yaml:"smtp_host"`
	SmtpEmail  string `mapstructure:"smtp_email" json:"smtp_email" yaml:"smtp_email"`
	SmtPass    string `mapstructure:"smtp_pass" json:"smtp_pass" yaml:"smtp_pass"`
}
