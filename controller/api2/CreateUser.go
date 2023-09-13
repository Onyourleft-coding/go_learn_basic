package api2

import (
	"gin_blog/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	//database.InsertUserWithTable()
	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
		"code":    http.StatusOK,
		"data":    0,
	})
}

func TakeUserInfo(c *gin.Context) {
	database.TakeUserInfo()
}
