package main

import (
	setting "back/pkg"
	"back/pkg/routers"
	"fmt"
)

func start() {
	router := routers.SetupRouter()

	err := router.Run(fmt.Sprintf(":%d", setting.HTTPPort))
	fmt.Println("text")
	if err != nil {
		return
	}
}

func main() {
	// 创建一个默认的路由引擎
	start()
}
