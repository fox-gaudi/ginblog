package main

import (
	"ginblog/app/models"
	"ginblog/routes"
)

func main() {
	// 初始化数据库
	models.ConnectDB()

	// 引入路由
	routes.InitRouter()
}
