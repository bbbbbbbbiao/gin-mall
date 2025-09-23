package v1

import (
	"gin-mall/app/common/request"
	"gin-mall/app/common/response"
	"gin-mall/app/serializer"
	"gin-mall/app/service"
	"github.com/gin-gonic/gin"
)

/**
 * @author: biao
 * @date: 2025/9/22 14:56
 * @code: 彼方尚有荣光在
 * @description: 商品Controller
 */

// 上传商品
func UploadProduct(c *gin.Context) {
	var produceInfo request.ProductInfo

	form, _ := c.MultipartForm()
	files := form.File["file"]

	id, _ := c.Get("id")

	if err := c.ShouldBind(&produceInfo); err != nil {
		response.BusinessFail(c, err.Error())
		return
	}

	err, product := service.ProductService.CreateProduct(id.(uint), produceInfo, files)

	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}

	response.Success(c, serializer.BuildProduct(product))

}
