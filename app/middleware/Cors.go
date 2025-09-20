package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

/**
 * @author: biao
 * @date: 2025/9/2 20:29
 * @code: 彼方尚有荣光在
 * @description: 处理跨域中间件
 */

func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowCredentials = true
	// 可以暴露什么东西
	config.ExposeHeaders = []string{"New-Token", "New-Expires-In", "Content-Disposition"}

	return cors.New(config)
}
