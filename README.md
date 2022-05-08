# GoWebDemo
Web applications developed using go's gin framework and gorm framework

使用go中的gin框架（web框架）和gorm框架（数据库）搭建的web简易web应用。

### 简介

整个应用分为三层，controller，service，repository。

controller层负责路由分发，分配页面不同路由逻辑的处理函数。

service层负责对应路由的处理，从controller层获取参数操纵repository数据层的curd

repository层为数据层，负责提供数据的CURD

### 运行

要运行程序需要在repository中新建文件db_data.go，文件中写入

```go
package repository

const (
	USERNAME = ""//登录mysql的用户名
	PASSWORD = ""//密码
	IP       = ""//数据库ip地址，本地的话则为127.0.0.1
	PORT     = ""//端口号，一般mysql默认端口3306
)

```

配置完成后命令行 go run server.go 即可运行