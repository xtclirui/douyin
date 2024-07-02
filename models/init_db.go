package models

import (
	"My_douyin/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	var err error

	arg := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
		config.Info.DB.Username, config.Info.DB.Password, config.Info.DB.Host,
		config.Info.DB.Port, config.Info.DB.Database, config.Info.DB.Charset,
		config.Info.DB.ParseTime, config.Info.DB.Loc)
	log.Println(arg)
	fmt.Println(arg)

	sql := mysql.Open(arg)
	DB, err = gorm.Open(sql, &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&UserInfo{}, &Video{}, &Comment{}, &UserLogin{})
	if err != nil {
		panic(err)
	}
}
