package main

import (
	"anroid/dao"
	"anroid/routers"
	"github.com/gin-contrib/cors"
)

func main() {
	r := routers.SetUpRouter()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},              // 允许所有来源
		AllowMethods:     []string{"*"},              // 允许的请求方法
		AllowHeaders:     []string{"*"},              // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"}, // 允许暴露的响应头
		AllowCredentials: true,                       // 允许携带凭证（如Cookie）
	}))
	//创建连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close() // 程序退出关闭数据库连接
	r.Run(":9999")
}
