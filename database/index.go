package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func LinkDataBase() {
	username := "root"  //账号
	password := "root"  //密码
	host := "127.0.0.1" //数据库地址 可以是ip或者是域名
	port := 3306        //数据库端口
	Dbname := "ginblog" //数据库名
	timeout := "10s"    //连接超时，10秒
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//	连接MySql 获得DB类型实例，用于后面的数据库读写操作
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//高级配置 跳过默认事务
		//为了确保数据一致性，GORM会在事务执行写入操作（创建、更新、删除）
		//如果没有这方面的需求，可以在初始化时禁用它，可以获得60%的性能提升
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic("连接数据库失败，error" + err.Error())
	}
	//	连接成功
	DB = db
	fmt.Println("db连接成功", DB)
	CreateTableWithUser()
}

// 创建用户表

type User struct {
	ID         uint //默认使用ID作为主键
	Name       string
	Password   string
	Phone      string
	Status     int //1 正常 0关闭 -1 假删除
	CreateTime string
}

func CreateTableWithUser() {
	DB.Debug().AutoMigrate(&User{})
}

func InsertUserWithTable() {
	user := User{
		Name:       "Reese",
		Password:   "zhao1Gem",
		Phone:      "15915176666",
		Status:     1,
		CreateTime: time.DateTime,
	}
	DB.Debug().Create(user)
}

func TakeUserInfo() User {
	var user User
	//user = DB.Take(&user).RowsAffected
	return user
}
