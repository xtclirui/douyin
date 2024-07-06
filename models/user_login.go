package models

import (
	"errors"
	"sync"
)

type UserLogin struct {
	Id         int64 `gorm:"primary_key"`
	UserInfoId int64
	Username   string `gorm:"primary_key"`
	Password   string `gorm:"size:200;notnull"`
}

type UserLoginDAO struct{}

var (
	userLoginDAO  *UserLoginDAO
	userLoginOnce sync.Once
)

func NewUserLoginDao() *UserLoginDAO {
	userLoginOnce.Do(func() {
		userLoginDAO = &UserLoginDAO{}
	})
	return userLoginDAO
}

func (mu *UserLoginDAO) DirectQueryUserLogin(username, password string, login *UserLogin) error {
	if login == nil {
		return errors.New("login *UserLogin 为空")
	}
	DB.Where("username = ? AND password = ?", username, password).First(login)
	if login.Id == 0 {
		return errors.New("用户不存在或密码错误")
	}
	return nil
}
