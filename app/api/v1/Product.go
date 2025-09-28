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
 * @date: 2025/9/22 14:56
 * @code: 彼方尚有荣光在
 * @description: 商品Controller
 */

// 上传商品
func UploadProduct(c *gin.Context) {
	var productInfo request.ProductInfo

	form, _ := c.MultipartForm()
	files := form.File["file"]

	id, _ := c.Get("id")

	if err := c.ShouldBind(&productInfo); err != nil {
		response.BusinessFail(c, err.Error())
		return
	}

	err, product := service.ProductService.CreateProduct(id.(uint), productInfo, files)

	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}

	response.Success(c, serializer.BuildProduct(product))
}

// 商品列表展示
func ProductList(c *gin.Context) {
	var productInfo request.ProductInfo

	if err := c.ShouldBind(&productInfo); err != nil {
		response.BusinessFail(c, err.Error())
		return
	}

	err, productList, total := service.ProductService.ProductList(productInfo)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}

	response.Success(c, serializer.BuildDataList(uint(total), serializer.BuildProductList(productList)))
}

// 搜索商品
func SearchProduct(c *gin.Context) {
	var productInfo request.ProductInfo

	if err := c.ShouldBind(&productInfo); err != nil {
		response.BusinessFail(c, err.Error())
		return
	}

	err, productList, total := service.ProductService.ProductSearch(productInfo)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}

	response.Success(c, serializer.BuildDataList(uint(total), serializer.BuildProductList(productList)))
}

// 获取商品详细信息
func ProductInfoById(c *gin.Context) {
	var productInfo request.ProductInfo

	if err := c.ShouldBind(&productInfo); err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	id := c.Param("id")
	parseId, _ := strconv.ParseUint(id, 10, 0)

	err, product := service.ProductService.ProductInfoById(uint(parseId))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}

	response.Success(c, serializer.BuildProduct(product))
}

// 获取商品图片信息
func ProductImgInfoById(c *gin.Context) {
	var productImgInfo request.ProductImgInfo

	if err := c.ShouldBind(&productImgInfo); err != nil {
		response.BusinessFail(c, err.Error())
		return
	}

	id := c.Param("id")
	parseId, _ := strconv.ParseUint(id, 10, 0)
	err, productImgList := service.ProductService.ProductImgById(uint(parseId))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}

	response.Success(c, serializer.BuildDataList(uint(len(productImgList)), serializer.BuildProductImgList(productImgList)))
}

func Categories(c *gin.Context) {

	err, categoryList := service.ProductService.Categories()
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, serializer.BuildDataList(uint(len(categoryList)), serializer.BuildCategoryList(categoryList)))
}
