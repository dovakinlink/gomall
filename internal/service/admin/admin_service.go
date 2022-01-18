package admin

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gomall/internal/dao/pool"
	"gomall/internal/model"
	"gomall/pkg/global/log"
)

type adminService struct {
}

var AdminService *adminService

func (*adminService) CheckPassword(username string, password string) (bool, *model.Admin, error) {
	db := pool.GetDB()
	var admin model.Admin
	var count int64
	db = db.Where("username = ?", username).Find(&admin).Count(&count)
	if db.Error != nil {
		log.Logger.Error("database error ", log.String("unknown error", db.Error.Error()))
		return false, &admin, db.Error
	}
	if count == 0 {
		return false, &admin, errors.New("用户名或密码错误")
	}
	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	if err != nil {
		return false, &admin, errors.New("用户名或密码错误")
	}
	return true, &admin, nil
}
