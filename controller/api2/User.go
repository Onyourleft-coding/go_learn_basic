package api2

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginForm struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var form LoginForm
	if c.ShouldBind(&form) == nil {
	}
	if form.User == "user" && form.Password == "password" {
		c.JSON(http.StatusOK, gin.H{
			"message": "登录成功",
			"code":    http.StatusOK,
			"data":    form,
		})
	} else {
		//c.JSON(500)
	}
}
