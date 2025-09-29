package dao

import (
	"gin-mall/app/model"
	"gin-mall/global"
)

/**
 * @author: biao
 * @date: 2025/9/29 下午2:22
 * @code: 彼方尚有荣光在
 * @description: 地址数据访问层
 */

type addressDao struct {
}

var AddressDao = new(addressDao)

// AddAddress 添加地址
func (addressDao *addressDao) AddAddress(address *model.Address) (err error) {
	return global.App.DB.Create(&address).Error
}

// DeleteAddress 删除地址
func (addressDao *addressDao) DeleteAddress(userId uint, addressId uint) (address *model.Address, err error) {
	err = global.App.DB.Where("user_id = ? and id = ?", userId, addressId).Delete(&address).Error
	return
}

// ListAddress 获取地址列表
func (addressDao *addressDao) ListAddress(userId uint) (addressList []*model.Address, err error) {
	err = global.App.DB.Where("user_id = ?", userId).Find(&addressList).Error
	return
}

// UpdateAddress 更新地址
func (addressDao *addressDao) UpdateAddress(userId uint, addressId uint, address *model.Address) (err error) {
	return global.App.DB.Where("user_id = ? and id = ?", userId, addressId).Updates(&address).Error
}
