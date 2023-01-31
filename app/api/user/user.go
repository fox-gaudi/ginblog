package user

import (
	"ginblog/utils/error/code"
	"github.com/gin-gonic/gin"
)

// Show Path: app/api/user/user.go
/**
 * @api {get} /api/user/show 获取用户信息
 */
func Show(c *gin.Context) {

}

// Create Path: app/api/user/user.go
/**
 * @api {post} /api/user/create 创建用户
 */

type CreateParam struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func Create(c *gin.Context) {
	var createParam CreateParam
	if err := c.ShouldBind(&createParam); err != nil {
		c.JSON(code.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(code.StatusOK, "success")
}

// Update Path: app/api/user/user.go
/**
 * @api {post} /api/user/update 更新用户信息
 */
func Update(c *gin.Context) {

}

// Delete Path: app/api/user/user.go
/**
 * @api {delete} /api/user/delete 删除用户
 */
func Delete(c *gin.Context) {

}
