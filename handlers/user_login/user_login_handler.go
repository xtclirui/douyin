package user_login

import (
	"My_douyin/models"
	"My_douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserLoginResponse struct {
	models.BaseResponse
}

func UserLoginHandler(c *gin.Context) {
	username := c.Query("username")
	password, ok := c.Get("password")
	if username == "" || !ok {
		c.JSON(http.StatusOK, UserLoginResponse{
			models.BaseResponse{
				StatusCode: 1,
				StatusMsg:  "密码解析错误",
			},
		})
	}
	qloginresponse, err := service.QueryUserLogin(username, password.(string))
}
