package service

import (
	"gomall/internal/dao/pool"
	"gomall/internal/model"
	"gomall/pkg/common/response"
)

type productService struct {
}

var ProductService = new(productService)

// ListCategoryById 根据分类ID获取子分类列表
func (*productService) ListCategoryById(id int) (*[]response.CategoryResponse, error) {
	db := pool.GetDB()
	var resSlice []response.CategoryResponse
	rows, err := db.Raw("select id, parent_id, name, leaf_node from product_category where parent_id = ? and show_status = 1 order by sort asc", id).Rows()
	if err != nil {
		return &resSlice, err
	}
	defer rows.Close()
	for rows.Next() {
		var res response.CategoryResponse
		db.ScanRows(rows, &res)

		resSlice = append(resSlice, res)
	}
	return &resSlice, nil
}

func (*productService) ListProductByCategoryId(id int, pageSize int, pageNum int) (*[]model.Product, int, error) {
	db := pool.GetDB()
	var resSlice []model.Product
	//db.Raw("select * from product where product_category_id = ?", id).Count(&total)
	//rows, err := db.Raw("select * from product where product_category_id = ? limit ?, ?", id, (pageNum-1)*pageSize, pageSize).Rows()
	rows, totalPage, err := pool.Service.RawForPage("select * from product where product_category_id = ?", pageSize, pageNum, id)
	if err != nil {
		return &resSlice, totalPage, err
	}
	defer rows.Close()
	for rows.Next() {
		var res model.Product
		db.ScanRows(rows, &res)
		resSlice = append(resSlice, res)
	}
	return &resSlice, totalPage, nil
}
