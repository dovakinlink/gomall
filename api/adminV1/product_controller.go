package adminV1

import (
	"github.com/gin-gonic/gin"
	"gomall/internal/service/admin"
	"gomall/pkg/common/request"
	"gomall/pkg/common/response"
	"net/http"
)

// AddBrand 新增品牌信息
func AddBrand(c *gin.Context) {
	var addBrandRequest request.AddBrandRequest
	c.ShouldBindJSON(&addBrandRequest)

	ok, err := admin.BrandService.AddBrand(addBrandRequest.Name, addBrandRequest.Sort, addBrandRequest.Logo)
	if err != nil || !ok {
		c.JSON(http.StatusOK, response.FailMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg("新增品牌信息成功"))
	return
}

// AddProductCategory 新增商品分类
func AddProductCategory(c *gin.Context) {
	var addProductCategoryRequest request.AddProductCategoryRequest
	c.ShouldBindJSON(&addProductCategoryRequest)

	ok, err := admin.ProductService.AddProductCategory(addProductCategoryRequest.ParentId,
		addProductCategoryRequest.Name, addProductCategoryRequest.Level, addProductCategoryRequest.ShowStatus,
		addProductCategoryRequest.Sort, addProductCategoryRequest.Keywords, addProductCategoryRequest.Description)
	if err != nil || !ok {
		c.JSON(http.StatusOK, response.FailMsg(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg("新增产品分类成功"))
	return
}

// AddProduct 新增商品
func AddProduct(c *gin.Context) {

}
