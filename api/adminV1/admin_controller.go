package adminV1

import (
	"github.com/gin-gonic/gin"
	"gomall/internal/service/admin"
	"gomall/pkg/common/response"
	"gomall/pkg/global/token"
	"net/http"
)

// Login 管理员登陆
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	ok, admin, err := admin.AdminService.CheckPassword(username, password)
	if err != nil {
		c.JSON(http.StatusOK, response.FailMsg(err.Error()))
		return
	}
	if !ok {
		c.JSON(http.StatusOK, response.FailMsg("登陆失败"))
		return
	}
	token, err := token.GenerateToken(int64(admin.Id), admin.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailMsg("service internal error"))
		return
	}
	c.JSON(http.StatusOK, response.SuccessMsg(token))
}
