package main

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// type User struct {
// 	ID   int    `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
// 	Name string `json:"name" gorm:"column:name"`
// 	Age  int    `json:"age" gorm:"column:age"`
// }

type User struct {
	UUID uuid.UUID `gorm:"type:char(36);column:uuid"`
	Name string    `gorm:"column:name"`
	Age  int       `gorm:"column:age"`
	Role string    `gorm:"column:role"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UUID = uuid.New()

	if u.Role == "admin" {
		return errors.New("invalid role")
	}
	return
}

func CreateUser() {
	GL_DB.AutoMigrate(&User{})
	// //插入一条记录
	// user4 := User1{}
	// GL_DB.Create(&user4)
	// fmt.Println("插入数据成功:", user4.ID)

	//创建记录，所有字段都插入
	// 这里的字段会被自动填充
	// 1. ID: 0 ，  UpdatedAt: 当前时间
	// user3 := User1{
	// 	Name:  "Bob",
	// 	Age:   25,
	// 	Email: "bob@example.com",
	// }
	// GL_DB.Create(&user3)
	// fmt.Println("插入数据成功:", user3.ID)

	// // 创建记录并忽略字段 ，除了UpdatedAt字段以外的字段都插入
	// user2 := User1{
	// 	Name:  "Alice",
	// 	Email: "alice@example.com",
	// 	Age:   28,
	// }
	// GL_DB.Select("*").Omit("UpdatedAt").Create(&user2)
	// fmt.Println("插入数据成功:", user2.ID)

	// // 创建一个用户实例
	// user := User1{
	// 	Name:  "John Doe",
	// 	Age:   30,
	// 	Email: "alice@example.com",
	// }
	// // 插入指定的字段 只插入Name字段
	// // 这里的字段会被自动填充
	// // 1. ID: 0 ，  UpdatedAt: 当前时间
	// GL_DB.Select("Name").Create(&user)
	// fmt.Println("插入数据成功:", user.ID)

	// 	//插入一条记录
	//    //指针创建一条记录
	// 	user := &User{Name: "Jinzhu", Age: 17}
	// 	GL_DB.Create(user) // 直接使用指针创建
	// 	fmt.Printf("指针用户: %+v\n", user)
	// // 通过结构体创建
	// structuser := User{Name: "Jinzhu", Age: 18}
	// GL_DB.Model(&User{}).Create(&structuser) // 使用Model指定模型
	// fmt.Printf("结构体用户: %+v\n", structuser)

	// // 通过 map 创建
	// mapuser := map[string]interface{}{

	// 	"age":  19,
	// 	"name": "Jinzhu1",
	// }
	// GL_DB.Model(&User{}).Create(&mapuser) // 使用Model指定模型
	// fmt.Printf("Map用户: %v\n", mapuser)

	//
	// // 通过指针切片创建
	// zzslice := []*User{
	// 	{Name: "jinzhu1", Age: 20},
	// 	{Name: "jinzhu2", Age: 20},
	// 	{Name: "jinzhu3", Age: 20},
	// }
	// GL_DB.Table("users").Create(&zzslice) // 通过指定表名创建
	// for i, user := range zzslice {
	// 	fmt.Printf("指针切片用户 %d: %+v\n", i+1, user)
	// }

	// // 通过结构体切片创建
	// structslice := []User{
	// 	{Name: "jinzhu1", Age: 21},
	// 	{Name: "jinzhu2", Age: 21},
	// 	{Name: "jinzhu3", Age: 21},
	// }
	// GL_DB.Table("users").Create(&structslice) // 通过指定表名创建
	// for i, user := range structslice {
	// 	fmt.Printf("结构体切片用户 %d: %+v\n", i+1, user)
	// }

	// // 通过 map 切片创建
	// mapslice := []map[string]interface{}{
	// 	{"name": "jinzhu1", "age": 22},
	// 	{"name": "jinzhu2", "age": 22},
	// 	{"name": "jinzhu3", "age": 22},
	// }
	// GL_DB.Table("users").Create(&mapslice) // 通过指定表名创建
	// for i, user := range mapslice {
	// 	fmt.Printf("Map切片用户 %d: %v\n", i+1, user)
	// }

	adminuser := User{Name: "zhangsan", Age: 18, Role: "admin"}
	if err := GL_DB.Create(&adminuser).Error; err != nil {
		fmt.Println("创建管理员用户失败:", err)
	} else {
		fmt.Println("创建管理用户成功")
	}

	commonuser := GL_DB.Create(&User{Name: "lisi", Age: 20, Role: "user"})
	if commonuser.Error != nil {
		fmt.Println("创建普通用户失败:", commonuser.Error)
	} else {
		fmt.Println("创建普通用户成功")
	}

}
