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
	var ids = c.DefaultQuery("id", "0")
	var pageSizeS = c.DefaultQuery("pageSize", "10")
	var pageNumS = c.DefaultQuery("pageNum", "1")
	id, err := strconv.Atoi(ids)
	pageSize, err := strconv.Atoi(pageSizeS)
	pageNum, err := strconv.Atoi(pageNumS)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.FailMsg("参数错误"))
		return
	}
	resSlice, totalPage, err := service.ProductService.ListProductByCategoryId(id, pageSize, pageNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailMsg("service internal error"))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsgForPages(resSlice, pageSize, pageNum, totalPage))
	return
}
