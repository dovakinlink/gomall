package v1

import (
	"github.com/gin-gonic/gin"
	"gomall/internal/service"
	"gomall/pkg/common/response"
	"net/http"
	"strconv"
)

// ListCategoryById 根据商品分类ID获取子分类列表
func ListCategoryById(c *gin.Context) {
	var ids = c.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.FailMsg("参数错误"))
		return
	}

	resSlice, err := service.ProductService.ListCategoryById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailMsg("service internal error"))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg(resSlice))
	return
}

// ListProductByCategoryId 根据分类ID获取该分类下的所有商品列表
func ListProductByCategoryId(c *gin.Context) {
	var ids = c.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.FailMsg("参数错误"))
		return
	}
	resSlice, err := service.ProductService.ListProductByCategoryId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailMsg("service internal error"))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg(resSlice))
	return
}
