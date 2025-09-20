package model

import "gorm.io/gorm"

/**
 * @author: biao
 * @date: 2025/9/2 11:14
 * @code: 彼方尚有荣光在
 * @description: 公告信息
 */

type Notice struct {
	gorm.Model
	Text string `gorm:"type:text"`
}
