package api

import (
	"ginblog/app/api/user"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	routerUser := r.Group("api/user")
	{
		routerUser.GET("api/user/show", user.Show)
		routerUser.POST("api/user/create", user.Create)
		routerUser.POST("api/user/update", user.Update)
		routerUser.DELETE("api/user/delete", user.Delete)
	}
}
