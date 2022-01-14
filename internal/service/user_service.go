package service

import (
	"golang.org/x/crypto/bcrypt"
	"gomall/internal/dao/pool"
	"gomall/internal/model"
)

type userService struct {
}

var UserService = new(userService)

// CheckPassword 校验用户名密码
func (*userService) CheckPassword(username string, password string) (bool, model.Member, error) {
	db := pool.GetDB()
	var member model.Member
	var count int64

	db = db.Where("username = ?", username).Find(&member).Count(&count)
	if db.Error != nil {
		return false, member, db.Error
	}
	if count == 0 {
		return false, member, nil
	} else {
		// 校验密码
		err := bcrypt.CompareHashAndPassword([]byte(member.Password), []byte(password))
		if err != nil {
			return false, member, nil
		} else {
			return true, member, nil
		}
	}
}

// CheckMember 校验是否存在指定用户
func (*userService) CheckMember(id int64, username string) (bool, error) {
	db := pool.GetDB()
	var member model.Member
	var count int64
	db = db.Where("id = ? and username = ?", id, username).Find(&member).Count(&count)
	if db.Error != nil {
		return false, db.Error
	}
	if count == 0 {
		return false, nil
	} else {
		return true, nil
	}
}
