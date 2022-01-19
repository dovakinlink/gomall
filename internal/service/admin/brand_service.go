package admin

import (
	"errors"
	"gomall/internal/dao/pool"
	"gomall/internal/model"
)

type brandService struct {
}

var BrandService *brandService

// AddBrand 新增品牌信息
func (*brandService) AddBrand(name string, sort int, logo string) (bool, error) {
	if name == "" {
		return false, errors.New("品牌名称不能为空")
	}
	if sort == 0 {
		return false, errors.New("排序字段从1开始")
	}
	db := pool.GetDB()
	var count int64
	db.Where("sort = ?", sort).Count(&count)
	if count != 0 {
		return false, errors.New("排序字段重复")
	}

	brand := &model.Brand{
		Name: name,
		Sort: sort,
		Logo: logo,
	}
	db = db.Create(brand)
	if db.Error != nil {
		return false, db.Error
	}
	return true, nil
}
