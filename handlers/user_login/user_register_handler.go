package user_login

import (
	"My_douyin/models"
	"My_douyin/service/user_login"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 注册回复

type UserRegisterResponse struct {
	models.BaseResponse
	*user_login.QLoginResponse
}

func UserRegisterHandler(c *gin.Context) {
	username := c.Query("username")
	//password, ok := c.Get("password")
	_, ok := c.Get("password")
	if len(username) == 0 || ok == false {
		c.JSON(http.StatusOK, UserRegisterResponse{
			BaseResponse: models.BaseResponse{
				StatusCode: 1,
				StatusMsg:  "username or password error",
			},
		})
	}

}
