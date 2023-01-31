package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode string
	AppPort string

	DB         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	RedisHost string
	RedisPort string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误, 请检查文件路径")
	}

	LoadServer(file)
	LoadData(file)
	LoadRedis(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("APP_MODE").String()
	AppPort = file.Section("server").Key("APP_PORT").String()
}

func LoadData(file *ini.File) {
	DB = file.Section("database").Key("DB").String()
	DbHost = file.Section("database").Key("DB_HOST").String()
	DbPort = file.Section("database").Key("DB_PORT").String()
	DbUser = file.Section("database").Key("DB_USER").String()
	DbPassword = file.Section("database").Key("DB_PASSWORD").String()
	DbName = file.Section("database").Key("DB_NAME").String()
}

func LoadRedis(file *ini.File) {
	RedisHost = file.Section("redis").Key("REDIS_HOST").String()
	RedisPort = file.Section("redis").Key("REDIS_PORT").String()
}
