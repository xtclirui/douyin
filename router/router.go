package router

import (
	"My_douyin/handlers/user_login"
	"My_douyin/handlers/video"
	"My_douyin/middleware"
	"My_douyin/models"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	models.InitDB()
	r := gin.Default()

	r.Static("/static", "./static")
	// 分组
	bg := r.Group("/douyin")
	// 登录
	bg.POST("/user/login/", middleware.SHA, user_login.UserLoginHandler)
	// 注册
	bg.POST("/user/register/", middleware.SHA, user_login.UserRegisterHandler)
	// 发布视频
	bg.POST("/publish/action/", middleware.JWT, video.PublishVideoHandler)
	// 推流
	bg.GET("/feed/", video.FeedVideoListHandler)

	return r
}
