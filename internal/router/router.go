package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	v1 "gomall/api/v1"
	"gomall/internal/middleware"
	"gomall/pkg/common/response"
	"gomall/pkg/global/log"
	"net/http"
)

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	server := gin.Default()
	server.Use(Cors())
	server.Use(Recovery)

	group := server.Group("")
	{
		group.POST("/user/register", v1.Register)                                             // 注册
		group.POST("/user/login", v1.Login)                                                   // 登陆
		group.GET("/user/token", middleware.TokenCheck(), v1.TestToken)                       // token验证测试
		group.GET("/product/category/list/:id", middleware.TokenCheck(), v1.ListCategoryById) // 根据商品分类ID获取子分类列表
	}
	return server
}

// Cors 跨域处理
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") // 请求头部
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		// 允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Logger.Error("HttpError", zap.Any("HttpError", err))
			}
		}()

		c.Next()
	}
}

func Recovery(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Logger.Error("gin catch error: ", log.Any("gin catch error: ", r))
			c.JSON(http.StatusOK, response.FailMsg("system internal error"))
		}
	}()
	c.Next()
}
