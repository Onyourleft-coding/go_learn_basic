package api1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//请求参数

// 查询参数Query

func Query(c *gin.Context) {
	fmt.Println(c.Query("user"))
	query := c.Query("user")
	fmt.Println(c.GetQuery("user"))
	fmt.Println(c.QueryArray("user")) //拿到多个相同的查询参数
	fmt.Println(c.DefaultQuery("addr", "广东省"))
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"data":    query,
		"message": "查询成功",
	})
}

// 动态参数

func Param(c *gin.Context) {
	fmt.Println(c.Param("user_id"))
	fmt.Println(c.Param("book_id"))
	userId := c.Param("user_id")
	bookId := c.Param("book_id")
	type Data struct {
		userId string
		bookId string
	}

	data := Data{userId, bookId}
	fmt.Println("data", data)
	c.JSON(http.StatusOK,
		gin.H{
			"code":    http.StatusOK,
			"data":    data,
			"message": "查询成功",
		})
}

//表单PostForm 可以接受multipart/form-data; 和application/x-www-form-urlencoded

func PostForm(c *gin.Context) {
	fmt.Println(c.PostForm("name"))
	fmt.Println(c.PostFormArray("name"))
	fmt.Println(c.DefaultPostForm("addr", "深圳市")) //如果用户没传就使用默认值
	forms, err := c.MultipartForm()
	fmt.Println("forms", forms, err)
	c.JSON(http.StatusOK,
		gin.H{
			"message": "查询成功",
			"data":    forms,
			"code":    http.StatusOK,
		})
}

//原始参数 GetRawData

func Raw(c *gin.Context) {
	body, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	switch contentType {
	case "application/json":

		//	json结构到结构体
		type User struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		var user User
		err := json.Unmarshal(body, &user)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(user)
		c.JSON(http.StatusOK, gin.H{
			"message": "请求成功",
			"code":    http.StatusOK,
			"data":    user,
		})
	}
}

//封装一个解析json到结构体上的函数

func BindJson(c *gin.Context, obj any) (err error) {
	body, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	switch contentType {
	case "application/json":
		err = json.Unmarshal(body, &obj)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

//四大请求方式 Restful风格是指网络应用中就是资源定位和资源操作的风格。不是标准也不是协议
/*
GET：从服务器取出资源（一项或多项）
POST:在服务器新建一个资源
PUT:在服务器更新资源（客户端提供完整资源数据）
PATCH：在服务器更新资源（客户端提供需要修改的资源数据）
DELETE：从服务器删除资源
*/

/*
//以文字资源为例
GET /articles 文章列表
GET /articles/:id 文章详情
POST /articles 添加文章
PUT /articles/:id 修改某一篇文章
DELETE /articles/:id 删除某一篇文章
*/

type ArticleModel struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ResponseResult struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

//文章列表页面

func GetList(c *gin.Context) {
	//	包含搜索，分页
	articleList := []ArticleModel{
		{1, "Go语言入门1", "这篇文章是《GO》语言入门1"},
		{2, "Go语言入门2", "这篇文章是《GO》语言入门2"},
		{3, "Go语言入门3", "这篇文章是《GO》语言入门3"},
		{4, "Go语言入门4", "这篇文章是《GO》语言入门4"},
	}
	c.JSON(http.StatusOK, ResponseResult{http.StatusOK, articleList, "成功"})
}

//文章详情

func GetDetail(c *gin.Context) {
	//获取params中的id
	//fmt.Println(c.Param("id"))
	param := c.Query("id")
	articleList := []ArticleModel{
		{1, "Go语言入门1", "这篇文章是《GO》语言入门1"},
	}
	fmt.Println("param", param)
	//c.JSON(http.StatusOK, gin.H{
	//	"data":    articleList,
	//	"code":    http.StatusOK,
	//	"message": "查询成功",
	//})
	c.JSON(http.StatusOK, ResponseResult{
		http.StatusOK,
		articleList,
		"查询成功",
	})
}

//编辑文章

func UpdateArticle(c *gin.Context) {

}
