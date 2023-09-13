package api2

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Data struct {
	UserName  string
	userPhone int
	city      string
}

func GetUerInfo(c *gin.Context) {
	data := Data{
		"reese",
		15915173646,
		"shenzhen",
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "查询成功",
		"code":    http.StatusOK,
		"data":    data,
	})
}
