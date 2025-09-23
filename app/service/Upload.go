package service

import (
	"gin-mall/global"
	"gin-mall/utils"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

/**
 * @author: biao
 * @date: 2025/9/17 21:51
 * @code: 彼方尚有荣光在
 * @description: 上传服务
 */

type uploadService struct {
}

var UploadService = new(uploadService)

// 上传至本地
func (u *uploadService) UploadToLocalStatic(id uint, file multipart.File, item string) (err error, filePath string) {
	bId := strconv.Itoa(int(id))
	var basePath string
	nowUnix := time.Now().Unix()

	if item == "avatar" {
		basePath = "." + global.App.Config.Path.AvatarPath + "user_" + bId
	} else if item == "product" {
		basePath = "." + global.App.Config.Path.ProductPath + "boss_" + bId
	}

	if exist, _ := utils.PathExists(basePath); !exist {
		utils.CreateDir(basePath)
	}

	content, err := io.ReadAll(file) // 内容的类型为：[]byte
	avatarPath := basePath + "/" + strconv.FormatInt(nowUnix, 10) + ".jpg"

	if err != nil {
		return err, ""
	}

	err = os.WriteFile(avatarPath, content, 0666)

	if err != nil {
		return err, ""
	}

	if item == "avatar" {
		filePath = "user_" + bId + "/" + strconv.FormatInt(nowUnix, 10) + ".jpg"
	} else if item == "product" {
		filePath = "product_" + bId + "/" + strconv.FormatInt(nowUnix, 10) + ".jpg"
	}
	return nil, filePath
}

// 上传头像至本地
func (u *uploadService) UploadAvatarToLocalStatic(id uint, file multipart.File) (err error, filePath string) {
	return u.UploadToLocalStatic(id, file, "avatar")
}

// 上传商品至本地
func (u *uploadService) UploadProductToLocalStatic(id uint, file multipart.File) (err error, filePath string) {
	return u.UploadToLocalStatic(id, file, "product")
}
