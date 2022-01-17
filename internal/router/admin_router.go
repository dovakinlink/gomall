package router

import (
	"github.com/gin-gonic/gin"
	v1 "gomall/api/v1"
)

func NewAdminRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	server := gin.Default()
	server.Use(Cors())
	server.Use(Recovery)

	group := server.Group("")
	{
		group.POST("/user/register", v1.Register) // 注册
	}
	return server
}
