package database

import (
	"encoding/json"
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
	ID         uint   `json:"id"` //默认使用ID作为主键
	Name       string `json:"name" gorm:"size:12"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Status     int    `json:"status"` //1 正常 0关闭 -1 假删除
	CreateTime string `json:"createTime"`
}

func CreateTableWithUser() {
	DB.Debug().AutoMigrate(&User{})
}

func InsertUserWithTable() *User {
	user := User{
		Name:       "Reese",
		Password:   "zhao1Gem",
		Phone:      "15915176666",
		Status:     1,
		CreateTime: time.DateTime,
	}
	DB.Create(&user)
	return &user
}

func BatchInsertUsersWithTable() {
	var studentList []User
	for i := 0; i < 10; i++ {
		studentList = append(studentList, User{
			Name:       fmt.Sprintf("Reese%d", i+1),
			Password:   "zhao1Gem",
			Phone:      "15915176666",
			Status:     1,
			CreateTime: time.DateTime,
		})
	}
	fmt.Println("studentList", studentList)
	DB.Create(&studentList)
}

func TakeUserInfo() User {
	var user User
	DB.Take(&user)
	fmt.Println("user", user)
	return user
}

func TakeFirstUserInfo() User {
	var user User
	DB.First(&user)
	return user
}

func TakeLastUserInfo() User {
	var user User
	DB.Last(&user)
	return user
}

//根据主键查询 take的第二个参数，默认会根据主键查询，可以是数字、字符串

func QueryUserInfo(id string) User {
	fmt.Println("id", id)
	var user User
	user = User{} // 重新赋值
	err := DB.Take(&user, id).Error
	switch err {
	case gorm.ErrRecordNotFound:
		fmt.Println("没有找到", err)
	default:
		fmt.Println("sql错误")
	}

	fmt.Println("user", user)
	return user
}

func QueryUserInfo2(name string) User {
	fmt.Println("name", name)
	var user User
	//使用？作为占位符，将查询的内容放入? 可以有效防止sql注入
	//相当于 SELECT * FROM `students` WHERE name = '机器人27号' LIMIT 1
	DB.Take(&user, "name = ?", name)
	return user
}

func QueryTargetTotal(name string) int64 {
	fmt.Println("name", name)
	var user User
	count := DB.Find(&user, "name = ?", name).RowsAffected
	fmt.Println("count", count)
	return count
}
func QueryUserList() []User {
	var userList []User
	DB.Find(&userList, []int{5, 6, 7, 8, 9})
	//DB.Find(&userList,5,6,7,8,9) //一样的
	fmt.Println("userList", userList)
	return userList
}

func QueryUserListByName() []User {
	var userList []User
	DB.Find(&userList, "name in ?", []string{"Reese", "Reese1", "Reese2"})
	fmt.Println("userList", userList)
	return userList
}

func UpdateUserInfo(id string, name string) User {
	var user User
	DB.Take(&user, id)
	user.Name = name
	//全字段更新
	DB.Save(&user)
	return user
}

func SelectUpdateUserInfo(id string, name string) User {
	var user User
	DB.Take(&user, id)
	user.Name = name
	//全字段更新
	DB.Select("name").Save(&user)
	return user
}

func BatchUpdatePassword(status string, password string) []User {
	var userList []User
	fmt.Println("userList", &userList)

	DB.Find(&userList, "status = ?", status).Update("password", password)
	for _, user := range userList {
		data, _ := json.Marshal(user)
		fmt.Println("string(data)", string(data))
	}

	//DB.Model(&User{}).Where("status = ?", status).Update("password", password)
	//不过这种方式没有返回值
	return userList
}
