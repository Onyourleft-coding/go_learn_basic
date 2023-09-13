package api1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloWorld() {
	router := gin.Default()
	//这是默认的服务器，使用gin的Default方法创建一个路由Handler
	router.GET("/api/info", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello World!")
		//通过Http方法绑定路由规则和路由函数。 200表示正常响应==>http.StatusOK
	})

	//返回txt
	router.GET("/api/txt", func(context *gin.Context) {
		context.String(http.StatusOK, "返回txt")
	})

	//返回json
	router.GET("/api/json", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	//返回xml
	router.GET("/api/xml", func(context *gin.Context) {
		context.XML(http.StatusOK, gin.H{"user": "wuwei", "message": "hey", "status": http.StatusOK})
	})

	//yaml
	router.GET("/api/yaml", func(context *gin.Context) {
		context.YAML(http.StatusOK, gin.H{"user": "wuwei", "message": "hey", "status": http.StatusOK})
	})

	//返回html
	//先要使用loadHTMLGlob()或者LoadHTMLFiles()方法加载模板文件
	//router.LoadHTMLGlob("gin框架/templates/*")
	//router.GET("/api/html", func(context *gin.Context) {
	//	context.HTML(http.StatusOK, "index.html", gin.H{
	//		"title": "Main website",
	//	})
	//})

	//	返回结构体转json
	router.GET("/api/moreJson", func(context *gin.Context) {
		type Msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg := Msg{"wuwei", "hey", 24}
		// 注意 msg.Name 变成了 "user" 字段
		// 以下方式都会输出 :   {"user": "wuwei", "Message": "hey", "Number": 123}
		context.JSON(http.StatusOK, msg)
	})

	//重定向
	router.GET("/api/redirect", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})

	//查询参数query
	router.GET("/api/getQuery", func(context *gin.Context) {
		fmt.Println(context.Query("user"))
		params := context.Query("user")
		fmt.Println(context.GetQuery("user"))
		fmt.Println(context.QueryArray("user")) //拿到多个相同的查询参数
		fmt.Println(context.DefaultQuery("addr", "广东省"))
		context.JSON(http.StatusOK, gin.H{
			"params": params,
		})
	})

	//动态参数param

	router.Run(":8080")
	//	启动路由的Run方法监听端口
	//http.ListenAndServe(":8000", router)
}
