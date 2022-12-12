package routes

import (
	"ginblog/app/api/user"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	routerUser := r.Group("api/user")
	{
		routerUser.GET("api/user/show", user.Show)
		routerUser.POST("api/user/create", user.Create)
		routerUser.POST("api/user/update", user.Update)
		routerUser.DELETE("api/user/delete", user.Delete)
	}

	err := r.Run(utils.AppPort)
	if err != nil {
		panic("服务器启动失败")
	}
}
