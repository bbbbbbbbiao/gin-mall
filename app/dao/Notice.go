package dao

import (
	"gin-mall/app/model"
	"gin-mall/global"
)

/**
 * @author: biao
 * @date: 2025/9/20 13:53
 * @code: 彼方尚有荣光在
 * @description: 操作Notice数据库
 */

type noticeDao struct {
}

var NoticeDao = new(noticeDao)

func (noticeDao *noticeDao) GetNoticeById(id uint) (err error, notice *model.Notice) {
	err = global.App.DB.Where("id = ?", id).First(&notice).Error
	return
}
