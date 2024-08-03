package main

import (
	"fmt"
	"sky-take-out-gin/config"
	"sky-take-out-gin/route"
	"sky-take-out-gin/utils/MySQL"
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

	err = route.SetupHTTPRoute()
	if err != nil {
		fmt.Println("初始化路由失败, err: ", err)
		return
	}
	fmt.Println("初始化路由成功---------------------------")
}
