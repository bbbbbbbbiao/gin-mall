package dao

import (
	"errors"
	"gin-mall/app/model"
	"gin-mall/global"
	"gorm.io/gorm"
)

/**
 * @author: biao
 * @date: 2025/9/4 21:12
 * @code: 彼方尚有荣光在
 * @description: 操作User数据库
 */

type userDao struct {
}

var UserDao = new(userDao)

func (userDao *userDao) FindUserByUserName(userName string) (user *model.User, exist bool, err error) {
	err = global.App.DB.Where("user_name = ?", userName).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return user, true, nil
}

func (userDao *userDao) CreateUser(user *model.User) error {
	return global.App.DB.Create(&user).Error
}

func (userDao *userDao) GetUserById(id uint) (user *model.User, err error) {
	err = global.App.DB.Where("id = ?", id).First(&user).Error
	return
}

func (userDao *userDao) UpdateUserInfo(user *model.User) error {
	return global.App.DB.Model(&model.User{}).Where("id = ?", user.ID).Updates(user).Error
}
