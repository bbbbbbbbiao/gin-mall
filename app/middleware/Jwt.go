package middleware

import (
	"gin-mall/app/common/response"
	"gin-mall/app/service"
	"gin-mall/global"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/**
 * @author: biao
 * @date: 2025/9/9 14:34
 * @code: 彼方尚有荣光在
 * @description: token鉴权中间件
 */

// 解析JWT
func ParseJWTAuth(guardName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			global.App.Log.Error("token 是空的")
			response.TokenFail(c)
			c.Abort()
			return
		}
		tokenStr = tokenStr[len(service.TokenType)+1:]

		// Token 解析校验
		token, err := jwt.ParseWithClaims(tokenStr, &service.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.App.Config.Jwt.Secret), nil
		})

		if err != nil {
			response.TokenFail(c)
			c.Abort()
			return
		}

		claims := token.Claims.(*service.Claims)
		if claims.StandardClaims.Issuer != guardName {
			response.TokenFail(c)
			c.Abort()
			return
		}

		c.Set("token", token)
		c.Set("id", claims.ID)
	}
}

// 解析Email的JWT
func ParseEmailJWTAuth(guardName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")

		if tokenStr == " " {
			global.App.Log.Error("EmailToken 是空的")
			response.EmailTokenFail(c)
			c.Abort()
			return
		}

		tokenStr = tokenStr[len(service.TokenType)+1:]

		// EmailToken 解析校验
		token, err := jwt.ParseWithClaims(tokenStr, &service.EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.App.Config.Jwt.Secret), nil
		})

		if err != nil {
			global.App.Log.Error("解析EmailToken失败，", zap.Any("err", err))
			response.EmailTokenFail(c)
			c.Abort()
			return
		}

		claims := token.Claims.(*service.EmailClaims)

		if claims.StandardClaims.Issuer != guardName {
			global.App.Log.Error("使用的guardName与EmailToken中的不一致：", zap.String("guardName", guardName))
			response.EmailTokenFail(c)
			c.Abort()
			return
		}

		c.Set("token", token)
		c.Set("id", claims.ID)
		c.Set("email", claims.Email)
		c.Set("password", claims.Password)
		c.Set("operationType", claims.OperationType)
	}
}
