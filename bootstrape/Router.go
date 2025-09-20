package bootstrape

import (
	"gin-mall/app/middleware"
	"gin-mall/global"
	"gin-mall/routes"
	"github.com/gin-gonic/gin"
)

/**
 * @author: biao
 * @date: 2025/9/2 20:18
 * @code: 彼方尚有荣光在
 * @description: 初始化gin和路由
 */

func SetRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())

	//加载静态文件路径
	//r.StaticFS("/static", http.Dir("./static"))
	r.Static("/static/", "./static")
	apiRouter := r.Group("/api/v1")
	routes.ApiV1Router(apiRouter)

	return r
}

func RunServe() {
	r := SetRouter()
	r.Run(":" + global.App.Config.Service.HttpPort)
}
