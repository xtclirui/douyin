package user_login

import (
	"My_douyin/middleware"
	"My_douyin/models"
	"errors"
)

type QRegisterResponse struct {
	Userid int64  `json:"userid"`
	Token  string `json:"token"`
}

type ServiceQRegisterResponse struct {
	username string
	password string
	data     *QRegisterResponse
	userid   int64
	token    string
}

func PostUserLogin(username, password string) (*QRegisterResponse, error) {
	postUserLoginRequest := NewPostUserLoginRequest(username, password)
	return postUserLoginRequest.Do()
}

func NewPostUserLoginRequest(username, password string) *ServiceQRegisterResponse {
	return &ServiceQRegisterResponse{username: username, password: password}
}

func (sp *ServiceQRegisterResponse) Do() (*QRegisterResponse, error) {
	if err := sp.checkFormat(); err != nil {
		return nil, err
	}
	if err := sp.updateData(); err != nil {
		return nil, err
	}
	if err := sp.fillData(); err != nil {
		return nil, err
	}
	return sp.data, nil
}

// 检查username 和 password
func (sp *ServiceQRegisterResponse) checkFormat() error {
	if sp.username == "" {
		return errors.New("空用户名")
	}
	if len(sp.username) > 100 {
		return errors.New("用户名长度超出限制")
	}
	if sp.password == "" {
		return errors.New("空密码为")
	}
	return nil
}

func (sp *ServiceQRegisterResponse) updateData() error {

	// 初始化models.UserInfo
	userLogin := models.UserLogin{Username: sp.username, Password: sp.password}
	userInfo := models.UserInfo{User: &userLogin, Name: sp.username}

	userloginDAO := models.NewUserLoginDao()
	if userloginDAO.UserExitByUSerName(sp.username) {
		return errors.New("用户名已存在")
	}

	// 更新数据库
	userInfoDAO := models.NewUserInfoDAO()
	err := userInfoDAO.AddUserInfo(&userInfo)
	if err != nil {
		return err
	}

	// 颁发token
	token, err := middleware.ReleaseToken(userLogin)
	if err != nil {
		return err
	}
	sp.token = token
	sp.userid = userInfo.Id
	return nil

}

func (sp *ServiceQRegisterResponse) fillData() error {
	sp.data = &QRegisterResponse{
		Userid: sp.userid,
		Token:  sp.token,
	}
	return nil
}
