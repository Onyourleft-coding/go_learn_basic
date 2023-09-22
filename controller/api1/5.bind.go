package api1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
)

//gin中的bind可以很方便的将前端传递来的数据与结构体进行参数绑定和参数校验

//参数绑定
//在使用这个功能的时候，需要给结构体加上Tag json form uri xml yaml

//must bind 不用，校验失败会改状态码
//should bind 可以绑定 json, query, param, yaml, xml

//ShouldBindJson

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func ShouldBindJson(c *gin.Context) {
	var userinfo UserInfo
	fmt.Println("userInfo", userinfo)
	err := c.ShouldBindJSON(&userinfo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "参数错误",
			"data":    "",
			"code":    "500",
		})
		return
	}
	c.JSON(200, userinfo)
}

//绑定查询参数 tag对应为form

type UserForm struct {
	Name string `json:"name" form:"name"`
	Age  int    `json:"age" form:"age"`
	Sex  string `json:"sex" form:"sex"`
}

func ShouldBindQuery(c *gin.Context) {
	var userForm UserForm
	err := c.ShouldBindQuery(&userForm)
	fmt.Println("err", err)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "参数错误",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    http.StatusOK,
		"message": "查询成功",
		"data":    userForm,
	})
}

//绑定动态参数 tar对应为uri

type UserForm2 struct {
	Name string `json:"name" form:"name" uri:"name"`
	Age  int    `json:"age" form:"age" uri:"age"`
	Sex  string `json:"sex" form:"sex" uri:"sex"`
}

func ShouldBindUri(c *gin.Context) {
	var userForm2 UserForm2
	err := c.ShouldBindUri(&userForm2)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"data":    "",
			"message": "错误参数",
			"code":    500,
		})
		return
	}
	c.JSON(http.StatusOK, userForm2)
}

/*
根据请求头中的content-type去自动绑定
form-data的参数也用这个，tag用form
默认的tag就是form
*/

func ShouldBind(c *gin.Context) {
	var userInfo UserInfo
	err := c.ShouldBind(&userInfo)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"data":    "",
			"message": "错误参数",
			"code":    http.StatusOK,
		})
		return
	}
	c.JSON(200, userInfo)
}

//bind绑定器
//需要使用参数验证功能，需要加binding tag
/*
 //不能为空，并且不能没有这个字段
required : 必填字段如 binding:"required"

针对字符串的长度
min 最小长度 如binding:"min=5"
max 最大长度 如binding:"max=12"
len 长度 如binding:"len=6"
针对数字的大小
eq 等于 如binding:"eq=3"
ne 不等于 如binding:"ne=12"
gt 大于 如binding:"gt=10"
gte 大于 等于 如binding:"gte=10"
lt 小于 如binding:"le=10"
lte 小于等于 如binding:"lte=10"

针对同级字段的
eqfield 等于其他字段的值 如PassWord string `binding:"eqfield:Password"`
nefield 不等于其他字段的值

- 忽略字段 如binding:"-"
*/
/*
//gin内置验证器
// 枚举  只能是red 或green
oneof=red green

// 字符串
contains=fengfeng  // 包含fengfeng的字符串
excludes // 不包含
startswith  // 字符串前缀
endswith  // 字符串后缀

// 数组
dive  // dive后面的验证就是针对数组中的每一个元素

// 网络验证
ip
ipv4
ipv6
uri
url
// uri 在于I(Identifier)是统一资源标示符，可以唯一标识一个资源。
// url 在于Locater，是统一资源定位符，提供找到该资源的确切路径

// 日期验证  1月2号下午3点4分5秒在2006年
datetime=2006-01-02

*/

// 自定义校验的错误信息
// 当验证不通过时，会给出错误的信息，但是原始的错误信息不太友好，不利于用户查看
// 只需要给一个结构体加msg的tag
type customUserInfo struct {
	Username string `json:"username" binding:"required" msg:"用户名不能为空"`
	Password string `json:"password" binding:"min=4,max=6" msg:"密码长度不能小于3大于6位"`
	Email    string `json:"email" binding:"email" msg:"邮箱地址格式不正确"`
}

/*
当出现错误时，就可以来获取出错字段上的msg
err : 这个参数为ShouldBindJson返回的错误信息
obj : 这个参数为绑定的结构体
还有一点要注意，validator这个包要引用v10版本，否则会报错
*/

//GetValidMsg 返回结构体中的msg参数

func GetValidMsg(err error, obj any) string {
	//使用的时候要传obj的指针
	getObj := reflect.TypeOf(obj)
	//将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		//断言成功
		for _, e := range errs {
			//	循环每一个错误信息
			//	根据错误字段名 获取结构体的具体字段
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return err.Error()
}
