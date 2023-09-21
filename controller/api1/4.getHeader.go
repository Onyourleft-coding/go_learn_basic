package api1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
请求头相关
*/

//请求头参数获取

func GetHeader(c *gin.Context) {
	//首字母不分大小写 单词和单词之间用 - 连接
	//用于获取一个请求头 这里的打印值是一致的
	fmt.Println("User-Agent", c.GetHeader("User-Agent"))
	fmt.Println("user-agent", c.GetHeader("user-agent"))
	fmt.Println("user-aGent", c.GetHeader("user-aGent"))
	fmt.Println("user-AGent", c.GetHeader("user-AGent"))

	//Header是一个普通的map[string][]string
	fmt.Println("c.Request.Header", c.Request.Header)
	//如果是使用Get方法或者.GetHeader，那么可以不用区分大小写，且返回第一个value

	fmt.Println("c.Request.Header.Get", c.Request.Header.Get("User-Agent"))
	fmt.Println("c.Request.Header", c.Request.Header["User-Agent"])
	//如果是用map的取值方式，请注意大小写问题
	fmt.Println("map方式获取请求头", c.Request.Header["user-agent"])

	//自定义请求头 用get方法也是免大小写
	fmt.Println(c.Request.Header.Get("Token"))
	fmt.Println(c.Request.Header.Get("token"))
	c.JSON(http.StatusOK, gin.H{
		"message": "获取成功",
		"data":    c.Request.Header.Get("token"),
		"code":    http.StatusOK,
	})
}

//响应头相关

//设置响应头

func SetHeader(c *gin.Context) {
	c.Header("Token", "dsadasdsadsadsadsa")
	c.Header("Content-Type", "application/text; charset=utf-8")
	c.JSON(http.StatusOK, gin.H{
		"message": "请求成功",
		"data":    "看看响应头",
		"code":    http.StatusOK,
	})
}
