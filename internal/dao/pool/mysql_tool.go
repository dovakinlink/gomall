package pool

import (
	"database/sql"
	"fmt"
	"gomall/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _db *gorm.DB
var dsn string

type service struct {
}

var Service = new(service)

func init() {
	username := config.GetConfig().MySQL.User     // 账号
	password := config.GetConfig().MySQL.Password // 密码
	host := config.GetConfig().MySQL.Host         // 数据库地址
	port := config.GetConfig().MySQL.Port         // 数据库端口
	Dbname := config.GetConfig().MySQL.Name       // 数据库名称
	timeout := "10s"                              // 连接超时时间

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//fmt.Printf("dsn: " + dsn)
	var err error
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("数据库连接失败，error= " + err.Error())
	}

	sqlDB, _ := _db.DB()

	// 设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) // 最大连接数
	sqlDB.SetMaxIdleConns(20)  // 最大空闲连接数
}

func GetDsn() string {
	return dsn
}

func GetDB() *gorm.DB {
	return _db
}

func (*service) RawForPage(sql string, pageSize int, pageNum int, values ...interface{}) (*sql.Rows, int, error) {
	var totalPage int64
	_db.Raw(sql, values).Count(&totalPage)
	rows, err := _db.Raw(sql+" limit ?,?", values, (pageNum-1)*pageSize, pageSize).Rows()
	if err != nil {
		return rows, int(totalPage), err
	}
	return rows, int(totalPage), nil
}
