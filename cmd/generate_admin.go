package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gomall/internal/dao/pool"
	"gomall/internal/model"
)

func main() {

	var username string
	var password string

	flag.StringVar(&username, "u", "", "username")
	flag.StringVar(&password, "p", "", "password")
	flag.Parse()
	if username == "" && password == "" {
		fmt.Println("请正确输入用户名密码")
		return
	}

	db := pool.GetDB()
	hashPw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return
	}
	admin := &model.Admin{
		Username: username,
		Password: string(hashPw),
	}
	db = db.Create(admin)
	if db.Error != nil {
		fmt.Println("未知错误")
		return
	}
	fmt.Println("创建成功")
}
