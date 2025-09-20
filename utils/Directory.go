package utils

import (
	"gin-mall/global"
	"go.uber.org/zap"
	"os"
)

/**
 * @author: biao
 * @date: 2025/9/1 20:44
 * @code: 彼方尚有荣光在
 * @description: 文件工具
 */

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

// 创建文件夹
func CreateDir(filePath string) bool {
	err := os.MkdirAll(filePath, 755)
	if err != nil {
		global.App.Log.Error("创建文件夹失败：", zap.Any("err", err))
		return false
	}
	return true
}
