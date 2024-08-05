package main

import (
	"fmt"
	"sky-take-out-gin/pkg/common/config"
	"sky-take-out-gin/pkg/common/database/MySQL"
	"sky-take-out-gin/pkg/common/database/Redis"
	"sky-take-out-gin/route"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		fmt.Println("初始化配置失败, err: ", err)
		return
	}
	fmt.Println("初始化配置成功----------------------------------")

	err = MySQL.InitDB()
	if err != nil {
		fmt.Println("初始化数据库失败, err: ", err)
		return
	}
	fmt.Println("初始化MySQL数据库成功---------------------------")

	err = Redis.InitRedis()
	if err != nil {
		fmt.Println("初始化Redis失败, err: ", err)
		return
	}
	fmt.Println("初始化Redis成功---------------------------")

	err = route.SetupHTTPRoute()
	if err != nil {
		fmt.Println("初始化路由失败, err: ", err)
		return
	}
	fmt.Println("初始化路由成功---------------------------")
}
