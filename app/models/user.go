package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID          uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id" comment:"用户ID"`
	Username    string         `gorm:"unique;not null" json:"username" comment:"用户名"`
	Password    string         `gorm:"not null" json:"password" comment:"密码"`
	Mobile      string         `json:"mobile" comment:"手机号码"`
	Avatar      string         `json:"avatar" comment:"头像"`
	NickName    string         `gorm:"column:nickname" json:"nick_name" comment:"昵称"`
	RealName    string         `gorm:"column:real_name" json:"real_name" comment:"姓名"`
	Email       *string        `gorm:"unique_index" json:"email" comment:"邮箱"`
	Age         uint8          `json:"age" comment:"年龄"`
	Birthday    *time.Time     `json:"birthday" comment:"生日"`
	ActivatedAt int64          `gorm:"default:null" json:"activated_at" comment:"会员激活时间"`
	CreatedAt   int64          `gorm:"autoCreateTime" json:"created_at" comment:"创建时间"`
	UpdatedAt   int64          `gorm:"autoCreateTime" json:"updated_at" comment:"更新时间"`
	DeletedAt   gorm.DeletedAt `gorm:"default:null" json:"deleted_at" comment:"删除时间"`
}
