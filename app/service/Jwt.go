package service

import (
	"gin-mall/global"
	"github.com/dgrijalva/jwt-go"
	"time"
)

/**
 * @author: biao
 * @date: 2025/9/8 11:06
 * @code: 彼方尚有荣光在
 * @description: JWT鉴权服务
 */

type jwtService struct {
}

var JwtService = new(jwtService)

const (
	TokenType    = "bearer"
	AppGuardName = "app"
)

type Claims struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Authority int    `josn:"authority"`
	jwt.StandardClaims
}

type EmailClaims struct {
	ID            uint   `json:"id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	OperationType uint   `json:"operation_type"`
	jwt.StandardClaims
}

type TokenOutPut struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expire_in"`
	TokenType   string `json:"token_type"`
}

// 签发Token
func (jwtService *jwtService) CreateToken(id uint, userName string, authority int, guardName string) (tokenData *TokenOutPut, err error, token *jwt.Token) {
	nowTimw := time.Now()
	claims := Claims{
		ID:        id,
		UserName:  userName,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: nowTimw.Unix() + global.App.Config.Jwt.JwtTtl,
			Issuer:    guardName, //用于在中间件中区分不同客户端颁发的token， 避免token 跨端使用
		},
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 使用Hash256进行加密

	tokenStr, err := token.SignedString([]byte(global.App.Config.Jwt.Secret)) //在对其进行一次签名

	tokenData = &TokenOutPut{
		AccessToken: tokenStr,
		ExpiresIn:   int(global.App.Config.Jwt.JwtTtl),
		TokenType:   TokenType,
	}
	return
}

// 签发EmailToken
func (jwtService *jwtService) CreateEamilToken(id uint, email string, password string, operationType uint, guardName string) (tokenData *TokenOutPut, err error, token *jwt.Token) {
	emailClaims := &EmailClaims{
		ID:            id,
		Email:         email,
		Password:      password,
		OperationType: operationType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + global.App.Config.Jwt.JwtTtl,
			Issuer:    guardName,
		},
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, emailClaims)
	tokenStr, err := token.SignedString([]byte(global.App.Config.Jwt.Secret))

	tokenData = &TokenOutPut{
		AccessToken: tokenStr,
		ExpiresIn:   int(global.App.Config.Jwt.JwtTtl),
		TokenType:   TokenType,
	}
	return
}
