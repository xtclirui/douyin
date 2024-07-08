package user_login

import (
	"My_douyin/models"
	"My_douyin/service/user_login"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 注册回复

type RegisterResponse struct {
	models.BaseResponse
	*user_login.QRegisterResponse
}

func UserRegisterHandler(c *gin.Context) {
	username := c.Query("username")
	password, _ := c.Get("password")

	registerResponse, err := user_login.PostUserLogin(username, password.(string))
	if err != nil {
		c.JSON(http.StatusOK, RegisterResponse{
			BaseResponse: models.BaseResponse{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK, RegisterResponse{
		BaseResponse: models.BaseResponse{
			StatusCode: 0,
		},
		QRegisterResponse: registerResponse,
	})
}
