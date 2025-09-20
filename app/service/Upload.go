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

// 上传头像服务至本地
func (u *uploadService) UploadAvatarToLocalStatic(id uint, file multipart.File) (err error, filePath string) {
	bId := strconv.Itoa(int(id))
	basePath := "." + global.App.Config.Path.AvatarPath + "user" + bId

	if exist, _ := utils.PathExists(basePath); !exist {
		utils.CreateDir(basePath)
	}

	content, err := io.ReadAll(file) // 内容的类型为：[]byte
	avatarPath := basePath + "/" + strconv.FormatInt(time.Now().Unix(), 10) + ".jpg"

	if err != nil {
		return err, ""
	}

	err = os.WriteFile(avatarPath, content, 0666)

	if err != nil {
		return err, ""
	}

	return nil, "user" + bId + "/" + strconv.FormatInt(time.Now().Unix(), 10) + ".jpg"
}
