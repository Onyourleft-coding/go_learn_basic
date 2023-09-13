package api1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin中的中间件必须是一个gin.HandlerFunc类型
func inderHandler(c *gin.Context) {
	fmt.Println("index.....")
	c.JSON(http.StatusOK, gin.H{
		"msg": "index",
	})
}

// 定义一个中间件
func m1(c *gin.Context) {
	fmt.Println("m1 in .....")
}

func handleRouter() {

	r := gin.Default()
	//	m1处于indexHandler函数前， 请求进来之后先走m1再走index
	r.GET("/api/middleIndex", m1, inderHandler)
	r.Run(":9899")
}
