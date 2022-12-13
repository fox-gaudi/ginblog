package routes

import (
	"ginblog/routes/admin"
	"ginblog/routes/api"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	// 后台路由
	admin.UserRouter(r)

	// 前端路由
	api.UserRouter(r)

	err := r.Run(utils.AppPort)
	if err != nil {
		panic("服务器启动失败")
	}
}
