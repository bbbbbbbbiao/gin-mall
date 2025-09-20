package config

/**
 * @author: biao
 * @date: 2025/9/1 16:08
 * @code: 彼方尚有荣光在
 * @description: 数据库配置结构体
 */

type Mysql struct {
	DB                  string `mapstructure:"db" json:"db" yaml:"db"`
	DbHost              string `mapstructure:"db_host" json:"db_host" yaml:"db_host"`
	DbPort              string `mapstructure:"db_port" json:"db_port" yaml:"db_port"`
	DbUser              string `mapstructure:"db_user" json:"db_user" yaml:"db_user"`
	DbPassword          string `mapstructure:"db_password" json:"db_password" yaml:"db_password"`
	DbName              string `mapstructure:"db_name" json:"db_name" yaml:"db_name"`
	LogMode             string `mapstructure:"log_mode" json:"log_mode" yaml:"log_mode"`
	EnableFileLogWriter bool   `mapstructure:"enable_file_log_writer" json:"enable_file_log_writer" yaml:"enable_file_log_writer"`
	LogFilename         string `mapstructure:"log_filename" json:"log_filename" yaml:"log_filename"`
}
