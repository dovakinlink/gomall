package main

import (
	"fmt"
	"github.com/gohouse/converter"
)

func main() {
	t2t := converter.NewTable2Struct()
	err := t2t.
		EnableJsonTag(true).
		PackageName("model").
		TagKey("gorm").
		RealNameMethod("TableName").
		SavePath("/Users/liuhuanchao/Downloads/model.go").
		Dsn("root:root453121@tcp(127.0.0.1:3306)/gomall?charset=utf8&parseTime=True&loc=Local&timeout=10s").
		Run()
	if err != nil {
		fmt.Println(err)
	}
}
