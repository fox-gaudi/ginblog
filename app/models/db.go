package models

import (
	"fmt"
	"ginblog/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB
var err error

func ConnectDB() {
	// Connect to the database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Shanghai", utils.DbHost, utils.DbUser, utils.DbPassword, utils.DbName, utils.DbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("数据库连接失败, 请检查数据库配置")
		panic(err)
	}

	fmt.Println("数据库连接成功")

	// Migrate the schema
	db.AutoMigrate(&User{})

	sqlDB, err := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Second * 10)

	// Close the database
	defer sqlDB.Close()
}
