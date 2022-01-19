package admin

import (
	"errors"
	"gomall/internal/dao/pool"
	"gomall/internal/model"
)

type productService struct {
}

var ProductService *productService

// AddProductCategory 新增商品分类
func (*productService) AddProductCategory(parentId int,
	name string, level int, showStatus int, sort int, keywords string, description string) (bool, error) {
	if name == "" {
		return false, errors.New("产品分类名称不能为空")
	}
	var count int64
	db := pool.GetDB()
	db.Where("id = ?", parentId).Count(&count)
	if count == 0 {
		return false, errors.New("产品父级ID不存在，请核对")
	}
	db.Where("sort = ? and parent_id = ?", sort, parentId).Count(&count)
	if count != 0 {
		return false, errors.New("排序字段重复")
	}
	productCategory := &model.Product_category{
		ParentId:    parentId,
		Name:        name,
		Level:       level,
		ShowStatus:  showStatus,
		Sort:        sort,
		Keywords:    keywords,
		Description: description,
		LeafNode:    0,
	}
	db = db.Create(productCategory)
	if db.Error != nil {
		return false, db.Error
	}
	return true, nil
}
