package main

import (
	"gin-mall/bootstrape"
	"gin-mall/global"
)

/**
 * @author: biao
 * @date: 2025/9/1 15:17
 * @code: 彼方尚有荣光在
 * @description: 主函数，程序入口
 */

func main() {
	bootstrape.InitializeConfig()
	global.App.Log = bootstrape.InitializeLog()

	//初始化数据库
	global.App.DB = bootstrape.InitializeDB()
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	//初始化验证器
	bootstrape.InitializeValidator()

	//启动服务器
	bootstrape.RunServe()
}
