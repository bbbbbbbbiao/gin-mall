package serializer

import "gin-mall/app/model"

/**
 * @author: biao
 * @date: 2025/9/29 下午2:39
 * @code: 彼方尚有荣光在
 * @description: 地址序列化器
 */

type Address struct {
	Id       uint   `json:"id"`
	UserId   uint   `json:"user_id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	CreateAt int64  `json:"create_at"`
}

func BuildAddress(address *model.Address) *Address {
	return &Address{
		Id:       address.ID,
		UserId:   address.UserID,
		Name:     address.Name,
		Phone:    address.Phone,
		Address:  address.Address,
		CreateAt: address.Model.CreatedAt.Unix(),
	}
}

func BuildAddressList(addresses []*model.Address) (addressList []*Address) {
	for _, address := range addresses {
		addressList = append(addressList, BuildAddress(address))
	}
	return
}
