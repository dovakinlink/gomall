package middleware

import (
	"github.com/gin-gonic/gin"
	"gomall/internal/service"
	"gomall/pkg/common/response"
	"gomall/pkg/global/token"
	"net/http"
)

// TokenCheck token校验
func TokenCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := c.GetHeader("Authorization")
		if t == "" {
			c.JSON(http.StatusNetworkAuthenticationRequired, response.FailMsg("请登录"))
			return
		}
		memberToken, err := token.ParseToken(t)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.FailMsg("token解析失败"))
			return
		}
		isExist, err := service.UserService.CheckMember(memberToken.ID, memberToken.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.FailMsg("service internal error"))
			return
		}
		if isExist {
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, response.FailMsg("鉴权失败"))
			return
		}

	}
}
