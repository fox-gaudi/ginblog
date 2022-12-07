package models

import (
	"database/sql"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID           uuid.UUID      `gorm:"type:uuid;primary_key;" json:"id" comment:"用户ID"`
	Name         string         `gorm:"type:varchar(20);not null" json:"name" comment:"姓名"`
	Password     string         `json:"password" comment:"密码"`
	Avatar       string         `json:"avatar" comment:"头像"`
	NickName     string         `gorm:"column:nickname" json:"nick_name" comment:"昵称"`
	Email        *string        `gorm:"unique_index" json:"email" comment:"邮箱"`
	Age          uint8          `json:"age" comment:"年龄"`
	Birthday     *time.Time     `json:"birthday" comment:"生日"`
	MemberNumber sql.NullString `gorm:"unique;not null" json:"member_number" comment:"会员编号"`
	ActivatedAt  int64          `gorm:"default:null" json:"activated_at" comment:"会员激活时间"`
	CreatedAt    int64          `gorm:"autoCreateTime" json:"created_at" comment:"创建时间"`
	UpdatedAt    int64          `gorm:"autoCreateTime" json:"updated_at" comment:"更新时间"`
	DeletedAt    gorm.DeletedAt `gorm:"default:null" json:"deleted_at" comment:"删除时间"`
}
