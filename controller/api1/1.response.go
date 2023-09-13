package api1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//返回字符串

func ResponseTxt(c *gin.Context) {
	c.String(http.StatusOK, "返回txt")
}

//返回json

func Response(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "查询成功",
		"data":    http.StatusOK,
		"code":    http.StatusOK,
	})
}

//返回结构体转json

func ResponseMoreJson(c *gin.Context) {
	//	可以在此处自定义一个结构体用于返回
	type Msg struct {
		Name    string `json:"user"`
		Message string
		Number  int
	}
	msg := Msg{"钟钟", "hey", 22}
	//	注意msg.Name变成了“user”字段
	//输出  {
	//	"user": "钟钟",
	//	"Message": "hey",
	//	"Number": 22
	//}
	c.JSON(http.StatusOK, msg)
}

//返回xml

func ResponseXml(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{
		"user":    "zhongzhong",
		"message": "hey xml",
		"status":  http.StatusOK,
	})
}

//返回yaml

func ResponseYaml(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{
		"user":    "zhongzhong",
		"message": "hey yaml",
		"status":  http.StatusOK,
	})
}

//返回html

func ResponseHtml(c *gin.Context) {
	//先要加载模板
	//	根据完整文件名渲染模板并传递参数
	c.HTML(http.StatusOK, "template1.html", gin.H{
		"title": "Main website",
	})
}

//返回重定向

func ResponseRedirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
}
