package main

import (
	"GoWeb/applications"
	"GoWeb/conf"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strconv"
)

func main() {
	// 创建一个echo实例
	// 不同于iris、gin 是直接使用New
	e := echo.New()

	// 注册中间件
	// 需要我们在入口文件手动注入基础中间件
	// iris、gin是又封装了一次 其实对我们来讲挺好的
	// 只是iris、gin命名封装的方法Default 实在没有语意
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 注册路由
	// 这里这几个框架的写法没啥区别
	e.GET("/", applications.Hello)

	// 启动服务
	// 基本同gin
	// 只是加了个返回值的日志
	p, _ := conf.ReadConf("./conf/settings.toml")

	port :=":"+strconv.Itoa(p.Port)
	e.Logger.Fatal(e.Start(port))
}

