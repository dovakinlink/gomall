package router

import (
	"github.com/gin-gonic/gin"
	"gomall/api/adminV1"
)

func NewAdminRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	server := gin.Default()
	server.Use(Cors())
	server.Use(Recovery)

	group := server.Group("")
	{
		group.POST("/admin/login", adminV1.Login) // 管理员登陆
	}
	return server
}
