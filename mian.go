package main

import (
	"My_douyin/config"
	"My_douyin/router"
	"fmt"
)

func main() {
	r := router.InitRouter()

	err := r.Run(fmt.Sprintf(":%d", config.Info.Port))
	if err != nil {
		return
	}
}
