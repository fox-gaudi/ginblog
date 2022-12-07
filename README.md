# Gin Blog

> It's just an exercise blog for beginners。

[![Go Report Card](https://goreportcard.com/badge/github.com/gin-gonic/gin)](https://goreportcard.com/report/github.com/gin-gonic/gin)[![GoDoc](https://pkg.go.dev/badge/github.com/gin-gonic/gin?status.svg)](https://pkg.go.dev/github.com/gin-gonic/gin?tab=doc)[![Release](https://img.shields.io/github/release/gin-gonic/gin.svg?style=flat-square)](https://github.com/gin-gonic/gin/releases)



项目代码用作练习。

![Gopher image](https://golang.org/doc/gopher/fiveyears.jpg)

## 使用指南

可以在本地开发或者使用docker

### 项目使用条件

Go 版本 1.19

```go
# 本地环境安装go
go run main.go
```

### 通过docker-compose 安装

本地需要安装docker 和 docker-compose。

获取项目&&进入项目根目录

```sh
# ssh
git clone git@github.com:fox-gaudi/ginblog.git && cd ginblog.git

# 或者https
git clone https://github.com/fox-gaudi/ginblog.git && cd ginblog.git
```

创建数据库所需要的目录:

```sh
# ginblog 根目录下 如果没有container_data就创建
mkdir container_data && cd container_data

# container_data
mkdir postgres
mkdir redis
```

运行项目:

```sh
# -d 在后台运行
docker-compose up
```

### 项目配置文件

```
# 在config目录下, 具体使用: https://ini.unknwon.cn/ 或 https://github.com/go-ini/ini
复制 config.ini.example 命名为 config.ini 
```

```ini
# config.ini
[server]
# mode: debug release test
APP_MODE = debug
APP_PORT = :80

# Database configuration
[database]
DB = postgres
DB_HOST = gin_blog_postgres // 使用本地的话, 直接用localhost, 文件默认docker-compose环境
DB_PORT = 5432
DB_USER = postgres
DB_PASSWORD = password
DB_NAME = postgres

[redis]
REDIS_HOST = gin_blog_redis
REDIS_PORT = 6379
```



## Authors 关于作者

* **Gaudi Fox** - Email : foxgaudi@gmail.com

## License 授权协议

这个项目 MIT 协议， 请点击 [LICENSE.md](LICENSE.md) 了解更多细节。
