package admin

import (
	"ginblog/app/admin/user"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	routerUser := r.Group("admin/user")
	{
		routerUser.GET("admin/user/index", user.Index)
	}
}
