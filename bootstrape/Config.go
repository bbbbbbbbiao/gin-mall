package bootstrape

import (
	"fmt"
	"gin-mall/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

/**
 * @author: biao
 * @date: 2025/9/1 16:45
 * @code: 彼方尚有荣光在
 * @description: 初始化配置文件，将配置文件的数据放到对应的结构体中
 */

func InitializeConfig() *viper.Viper {
	config := "config.yaml"

	if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
		config = configEnv
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed, err: %s \n", err))
	}

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)

		//重载配置
		if err := v.Unmarshal(&global.App.Config); err != nil {
			fmt.Println(err)
		}
	})

	// 将配置赋值给全局变量

	if err := v.Unmarshal(&global.App.Config); err != nil {
		fmt.Println(err)
	}

	return v
}
