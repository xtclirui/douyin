package user_login

import (
	"My_douyin/models"
	"My_douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 返回值

type UserLoginResponse struct {
	// 状态码
	// 状态码描述
	models.BaseResponse
	// 用户id
	// 用户鉴权token
	*service.QLoginResponse
}

func UserLoginHandler(c *gin.Context) {
	// 1. 校验参数的正确性，参数是否为空
	username := c.Query("username")
	password, _ := c.Get("password")
	//if username == "" || !ok {
	//	c.JSON(http.StatusOK, UserLoginResponse{
	//		BaseResponse: models.BaseResponse{
	//			StatusCode: 1,
	//			StatusMsg:  "密码解析错误",
	//		},
	//	})
	//}

	// 2. 通过用户名查询user实体，判断用户名是否存在
	// type QLoginResponse struct, err
	qloginresponse, err := service.QueryUserLogin(username, password.(string))

	// 用户不存在
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			BaseResponse: models.BaseResponse{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
	}

	// 3. 通过user实体，对密码进行加密后，与数据库中密码比对，是否一致
	// 用户存在,返回相应的id和token
	c.JSON(http.StatusOK, UserLoginResponse{
		BaseResponse: models.BaseResponse{
			StatusCode: 0,
		},
		QLoginResponse: qloginresponse,
	})

}
