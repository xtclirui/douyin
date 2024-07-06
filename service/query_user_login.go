package service

import (
	"My_douyin/middleware"
	"My_douyin/models"
	"errors"
)

// QLoginResponse 用户id 用户鉴权token
type QLoginResponse struct {
	Userid int64  `json:"userid"`
	Token  string `json:"token"`
}

type ServiceQLoginResponse struct {
	username string
	password string
	data     *QLoginResponse
	userid   int64
	token    string
}

func QueryUserLogin(username string, password string) (*QLoginResponse, error) {
	sQLoginResponse := NewServiceQloginresponse(username, password)
	return sQLoginResponse.Do()
}

func NewServiceQloginresponse(username string, password string) *ServiceQLoginResponse {
	return &ServiceQLoginResponse{username: username, password: password}
}

func (sq *ServiceQLoginResponse) Do() (*QLoginResponse, error) {
	if err := sq.checkFormat(); err != nil {
		return nil, err
	}
	if err := sq.findData(); err != nil {
		return nil, err
	}
	if err := sq.fillData(); err != nil {
		return nil, err
	}
	return sq.data, nil
}

// 检查username 和 password
func (sq *ServiceQLoginResponse) checkFormat() error {
	if sq.username == "" {
		return errors.New("空用户名")
	}
	if len(sq.username) > 100 {
		return errors.New("用户名长度超出限制")
	}
	if sq.password == "" {
		return errors.New("空密码为")
	}
	return nil
}

func (sq *ServiceQLoginResponse) findData() error {
	userloginDAO := models.NewUserLoginDao()
	var login models.UserLogin
	err := userloginDAO.DirectQueryUserLogin(sq.username, sq.password, &login)
	if err != nil {
		return err
	}
	sq.userid = login.UserInfoId

	// 生成token
	token, err := middleware.ReleaseToken(login)
	if err != nil {
		return err
	}
	sq.token = token
	return nil
}

func (sq *ServiceQLoginResponse) fillData() error {
	sq.data = &QLoginResponse{
		Userid: sq.userid,
		Token:  sq.token,
	}
	return nil
}
