package models

import (
	"database/sql"
	"fmt"
	"ginblog/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var err error

func ConnectDB() {
	// Connect to the database
	switch utils.DB {
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Shanghai", utils.DbHost, utils.DbUser, utils.DbPassword, utils.DbName, utils.DbPort)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		connectEroMsg(err)
		connectDB(db)
	case "mysql":
		panic("mysql is not supported yet")
	}

}

func connectDB(db *gorm.DB) {

	// Migrate the schema
	migrateEro := db.AutoMigrate(&User{})
	if migrateEro != nil {
		fmt.Println("数据库迁移失败")
		panic(err)
	}

	sqlDB, _ := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Second * 10)

	// Close the database
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
			panic("数据库关闭失败")
		}
	}(sqlDB)
}

func connectEroMsg(err error) {
	if err != nil {
		fmt.Println("数据库连接失败, 请检查数据库配置")
		panic(err)
	}
	fmt.Println("数据库连接成功")
}
