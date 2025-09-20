package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

/**
 * @author: biao
 * @date: 2025/9/4 21:28
 * @code: 彼方尚有荣光在
 * @description: 密码加密及密码验证
 */
func BcryptMake(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func BcryptMakeCheck(pwd []byte, hashedPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, pwd)
	if err != nil {
		return false
	}
	return true
}
