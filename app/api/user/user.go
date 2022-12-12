package user

import (
	"github.com/gin-gonic/gin"
)

// Show Path: app/api/user/user.go
/**
 * @api {get} /api/user/show 获取用户信息
 */
func Show(c *gin.Context) {
	// 先返回一个json数据
	c.JSON(200, gin.H{
		"name": c.Query("name"),
	})
}

// Create Path: app/api/user/user.go
/**
 * @api {post} /api/user/create 创建用户
 */
func Create(c *gin.Context) {

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
