package router

import (
	"gin_blog/controller/api1"
	"gin_blog/controller/api2"
	"github.com/gin-gonic/gin"
)

func StartRouter() *gin.Engine {
	// 整个路由的入口函数
	router := gin.Default()

	r := router.Group("/api") //接口请求前缀
	//响应
	{
		r.GET("/")
		r.GET("/responseTxt", api1.ResponseTxt)
		r.GET("/responseJson", api1.Response)
		r.GET("/responseMoreJson", api1.ResponseMoreJson)
		r.GET("/responseXml", api1.ResponseXml)
		r.GET("/responseYaml", api1.ResponseYaml)
		//先要加载模板
		//router.LoadHTMLGlob("gin_blog/templates/*")
		//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
		//r.GET("/responseHtml", api1.ResponseHtml)
		r.GET("/responseRedirect", api1.ResponseRedirect)
	}
	//请求
	{
		r.GET("/query", api1.Query)
		r.GET("/param/:user_id", api1.Param)
		r.GET("/param/:user_id/:book_id", api1.Param)
		r.POST("/postForm", api1.PostForm)
		r.POST("/raw", api1.Raw)
		//r.POST("/bindJson", api1.BindJson)
	}
	//Restful案例
	{
		r.GET("/getList", api1.GetList)
		r.GET("/getDetail", api1.GetDetail)
		r.POST("/createArticle", api1.CreateArticle)
		r.PUT("/updateArticle", api1.UpdateArticle)
		r.DELETE("/deleteArticle", api1.DeleteArticle)
	}
	//请求头相关
	{
		r.GET("/getHeader", api1.GetHeader)
	}
	//响应头相关
	{
		r.GET("/setHeader", api1.SetHeader)
	}
	//Bind
	{
		r.POST("/shouldBindJson", api1.ShouldBindJson)
		r.POST("/shouldBindQuery", api1.ShouldBindQuery)
		r.POST("/shouldBindUri", api1.ShouldBindUri)
		r.POST("/shouldBind", api1.ShouldBind)
	}
	v1 := router.Group("/v1")
	{
		v1.POST("/postForm")
	}
	v2 := router.Group("/v2").Use(api2.JwtTokenMiddleware)

	{
		v2.GET("/getUserInfo", api2.GetUerInfo)
		v2.POST("/CreateUser", api2.CreateUser)
		v2.POST("/batchInsertUser", api2.BatchInsertUsersWithTable)
		v2.GET("/takeUserInfo", api2.TakeUserInfo)
		v2.GET("/takeFirstUserInfo", api2.TakeFirstUserInfo)
		v2.GET("/takeLastUserInfo", api2.TakeLastUserInfo)
		v2.GET("/queryUserInfo", api2.QueryUserInfo)
		v2.GET("/queryUserInfo2", api2.QueryUserInfo2)
		v2.GET("/QueryTargetTotal", api2.QueryTargetTotal)
		v2.POST("/QueryUserList", api2.QueryUserList)
		v2.POST("/QueryUserListByName", api2.QueryUserListByName)
		v2.POST("/UpdateUserInfo", api2.UpdateUserInfo)
		v2.POST("/SelectUpdateUserInfo", api2.SelectUpdateUserInfo)
	}

	return router
}
