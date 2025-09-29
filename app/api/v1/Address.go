package v1

import (
	"gin-mall/app/common/request"
	"gin-mall/app/common/response"
	"gin-mall/app/serializer"
	"gin-mall/app/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

/**
 * @author: biao
 * @date: 2025/9/29 上午11:13
 * @code: 彼方尚有荣光在
 * @description: 地址Controller
 */

// AddAddress 添加地址
func AddAddress(c *gin.Context) {
	var addAddress request.Address

	if err := c.ShouldBind(&addAddress); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(addAddress, err))
		return
	}

	id, _ := c.Get("id")
	userId := uint(id.(int64))
	// 保存地址
	address, err := service.AddressService.AddAddress(userId, addAddress)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}

	response.Success(c, address)
}

func DeleteAddress(c *gin.Context) {
	var deleteAddress request.Address

	if err := c.ShouldBind(&deleteAddress); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(deleteAddress, err))
		return
	}

	id, _ := c.Get("id")
	userId := uint(id.(int64))
	idStr := c.Param("id")
	deleteId, _ := strconv.Atoi(idStr)
	// 删除地址
	address, err := service.AddressService.DeleteAddress(userId, uint(deleteId))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, address)
}

// ListAddress 获取地址列表
func ListAddress(c *gin.Context) {
	id, _ := c.Get("id")
	userId := uint(id.(int64))
	// 获取地址列表
	addressList, err := service.AddressService.ListAddress(userId)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, serializer.BuildDataList(uint(len(addressList)), serializer.BuildAddressList(addressList)))
}

// UpdateAddress 更新地址
func UpdateAddress(c *gin.Context) {
	var updateAddress request.Address

	if err := c.ShouldBind(&updateAddress); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(updateAddress, err))
		return
	}

	id, _ := c.Get("id")
	userId := uint(id.(int64))
	idStr := c.Param("id")
	updateId, _ := strconv.Atoi(idStr)

	// 更新地址
	address, err := service.AddressService.UpdateAddress(userId, uint(updateId), updateAddress)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, address)
}
