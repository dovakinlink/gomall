package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

// GoMall 万家医保配置信息
type GoMall struct {
	MySQL      MySQLConfig
	Log        LogConfig
	StaticPath PathConfig
	Token      TokenConfig
	Wechat     WechatConfig
}

// MySQLConfig 数据库配置信息
type MySQLConfig struct {
	Host        string
	Name        string
	Password    string
	Port        int
	TablePrefix string
	User        string
}

// LogConfig 日志配置信息
type LogConfig struct {
	Path  string
	Level string
}

// PathConfig 静态文件地址配置信息
type PathConfig struct {
	FilePath string
}

// TokenConfig jwt token配置
type TokenConfig struct {
	Key        string
	ExpireTime time.Duration
}

// WechatConfig 微信小程序配置信息
type WechatConfig struct {
	AppId     string
	AppSecret string
}

var c GoMall

func init() {
	// 设置文件名
	viper.SetConfigName("config")
	// 设置文件类型
	viper.SetConfigType("toml")
	// 设置文件路径，默认项目根路径
	viper.AddConfigPath(".")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	viper.Unmarshal(&c)
}

func GetConfig() GoMall {
	return c
}
