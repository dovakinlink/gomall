package v1

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gomall/internal/dao/pool"
	"gomall/internal/model"
	"gomall/internal/service"
	"gomall/pkg/common/request"
	"gomall/pkg/common/response"
	"gomall/pkg/global/log"
	"gomall/pkg/global/token"
	"net/http"
	"time"
)

// Register 注册
func Register(c *gin.Context) {
	var registerRequest request.RegisterRequest
	c.ShouldBindJSON(&registerRequest)

	var member *model.Member
	var count int64
	db := pool.GetDB()
	db.Where("username = ?", registerRequest.Username).Find(&member).Count(&count)

	// 不存在则注册
	if count == 0 {
		// 加密密码
		hashPw, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Logger.Error("service internal error: ", log.String("err: ", err.Error()))
			c.JSON(http.StatusOK, response.FailMsg(err.Error()))
			return
		}
		// 创建新的会员记录
		member = &model.Member{
			Username:   registerRequest.Username,
			Password:   string(hashPw),
			Nickname:   "",
			Phone:      "",
			Status:     0,
			CreateTime: time.Now().Format("2006-01-02 15:04:05"),
			UpdateTime: time.Now().Format("2006-01-02 15:04:05"),
		}
		db = db.Create(member)
		if db.Error != nil {
			log.Logger.Error("service internal error: ", log.String("err:", db.Error.Error()))
			c.JSON(http.StatusOK, response.FailMsg("service internal error"))
			return
		}
		c.JSON(http.StatusOK, response.SuccessMsg("注册成功"))
	} else {
		c.JSON(http.StatusOK, response.FailMsg("该用户已注册"))
	}
	return
}

// Login 登陆
func Login(c *gin.Context) {
	var loginRequest request.LoginRequest
	c.ShouldBindJSON(&loginRequest)
	var member model.Member
	isExist, member, err := service.UserService.CheckPassword(loginRequest.Username, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailMsg("service internal error"))
		return
	}
	if !isExist {
		c.JSON(http.StatusOK, response.FailMsg("用户名或密码错误"))
		return
	}
	t, err := token.GenerateToken(int64(member.Id), member.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailMsg("service internal error"))
		return
	}

	c.JSON(http.StatusOK, response.SuccessMsg(t))
	return
}

func TestToken(c *gin.Context) {
	c.JSON(http.StatusOK, response.SuccessMsg("TOKEN正确"))
}
