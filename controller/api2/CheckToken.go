package api2

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtTokenMiddleware(c *gin.Context) {
	//	获取请求头的token
	token := c.GetHeader("token")
	//	调用jwt函数
	if token == "1234" {
		//	校验通过
		c.Next()
		return
	}
	//	校验不通过
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusUnauthorized,
		"message": "权限校验失败",
		"data":    "",
	})
	c.Abort()
}
