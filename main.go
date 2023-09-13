package main

import (
	"gin_blog/database"
	"gin_blog/router"
)

func main() {
	database.LinkDataBase()
	//引入基础路由创建路由
	router := router.StartRouter()
	//端口号
	router.Run(":9898")
}
