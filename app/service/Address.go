package service

import (
	"gin-mall/app/common/request"
	"gin-mall/app/dao"
	"gin-mall/app/model"
	"gin-mall/global"
	"go.uber.org/zap"
)

/**
 * @author: biao
 * @date: 2025/9/29 下午2:15
 * @code: 彼方尚有荣光在
 * @description: 地址服务
 */

type addressService struct {
}

var AddressService = new(addressService)

// AddAddress 添加地址
func (a *addressService) AddAddress(userId uint, params request.Address) (address *model.Address, err error) {
	address = &model.Address{
		UserID:  userId,
		Name:    params.Name,
		Phone:   params.Phone,
		Address: params.Address,
	}
	err = dao.AddressDao.AddAddress(address)
	if err != nil {
		global.App.Log.Error("添加地址失败", zap.Error(err))
		return
	}

	return
}

// DeleteAddress 删除地址
func (a *addressService) DeleteAddress(userId uint, addressId uint) (address *model.Address, err error) {
	address, err = dao.AddressDao.DeleteAddress(userId, addressId)
	if err != nil {
		global.App.Log.Error("删除地址失败", zap.Error(err))
		return
	}

	return
}

// ListAddress 获取地址列表
func (a *addressService) ListAddress(userId uint) (addressList []*model.Address, err error) {
	addressList, err = dao.AddressDao.ListAddress(userId)
	if err != nil {
		global.App.Log.Error("获取地址列表失败", zap.Error(err))
		return
	}

	return
}

// UpdateAddress 更新地址
func (a *addressService) UpdateAddress(userId uint, addressId uint, params request.Address) (address *model.Address, err error) {
	address = &model.Address{
		UserID:  userId,
		Name:    params.Name,
		Phone:   params.Phone,
		Address: params.Address,
	}
	err = dao.AddressDao.UpdateAddress(userId, addressId, address)
	if err != nil {
		global.App.Log.Error("更新地址失败", zap.Error(err))
		return
	}
	return
}
