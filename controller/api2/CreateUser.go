package api2

import (
	"fmt"
	"gin_blog/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

//创建单个用户

func CreateUser(c *gin.Context) {
	fmt.Println(c.PostForm("id"))
	fmt.Println(c.MultipartForm())
	result := database.InsertUserWithTable()
	c.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
		"code":    http.StatusOK,
		"data":    result,
	})
}

//批量创建用户

func BatchInsertUsersWithTable(c *gin.Context) {
	database.BatchInsertUsersWithTable()
	c.JSON(http.StatusOK, gin.H{
		"message": "批量创建成功",
		"code":    http.StatusOK,
		"data":    "",
	})
}

//获取单条记录 一个用户

func TakeUserInfo(c *gin.Context) {
	result := database.TakeUserInfo()
	c.JSON(http.StatusOK, gin.H{
		"message": "查询成功",
		"code":    http.StatusOK,
		"data":    result,
	})
}

//获取第一个用户

func TakeFirstUserInfo(c *gin.Context) {
	result := database.TakeFirstUserInfo()
	c.JSON(http.StatusOK, gin.H{
		"message": "获取第一个用户成功",
		"code":    http.StatusOK,
		"data":    result,
	})
}

//获取最后一个用户

func TakeLastUserInfo(c *gin.Context) {
	result := database.TakeLastUserInfo()
	c.JSON(http.StatusOK, gin.H{
		"message": "获取最后一个用户成功",
		"code":    http.StatusOK,
		"data":    result,
	})
}

//根据主键查询

func QueryUserInfo(c *gin.Context) {
	id := c.Query("id")
	fmt.Println("query", id)
	result := database.QueryUserInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"message": "查询成功",
		"code":    http.StatusOK,
		"data":    result,
	})
}

//根据其他条件查询

func QueryUserInfo2(c *gin.Context) {
	name := c.Query("name")
	result := database.QueryUserInfo2(name)
	c.JSON(http.StatusOK, gin.H{
		"message": "查询成功",
		"code":    http.StatusOK,
		"data":    result,
	})
}

//获取查询结果

func QueryTargetTotal(c *gin.Context) {
	name := c.Query("name")
	result := database.QueryTargetTotal(name)
	c.JSON(http.StatusOK, gin.H{
		"message": "查询成功",
		"code":    http.StatusOK,
		"data":    result,
	})
}

//查询多条记录 根据主键列表查询

func QueryUserList(c *gin.Context) {
	result := database.QueryUserList()
	c.JSON(http.StatusOK, gin.H{
		"message": "查询成功",
		"code":    http.StatusOK,
		"data":    result,
	})
}

//查询多条记录 根据逐渐列表查询

func QueryUserListByName(c *gin.Context) {
	result := database.QueryUserListByName()
	c.JSON(http.StatusOK, gin.H{
		"message": "查询成功",
		"code":    http.StatusOK,
		"data":    result,
	})
}

//更新 更新的前提是先查询到记录
//save 保存所有字段 用户单个记录的全字段更新 它会保存所有字段 即使零值也会更新

func UpdateUserInfo(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	result := database.UpdateUserInfo(id, name)
	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
		"code":    http.StatusOK,
		"data":    result,
	})
}

// 更新指定字段

func SelectUpdateUserInfo(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	result := database.SelectUpdateUserInfo(id, name)
	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
		"code":    http.StatusOK,
		"data":    result,
	})
}

// 批量更新

func BatchUpdatePassword(c *gin.Context) {
	status := c.PostForm("status")
	password := c.PostForm("password")
	fmt.Println("name", status)
	fmt.Println("password", password)
	result := database.BatchUpdatePassword(status, password)
	c.JSON(http.StatusOK, gin.H{
		"message": "更新成功",
		"code":    0,
		"data":    result,
	})
}
