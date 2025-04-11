package main

// import (
// 	"fmt"
// 	"time"

// 	"gorm.io/gorm"
// )

// type User struct {
// 	Name string
// }

// type User1 struct {
// 	ID        int
// 	Name      string
// 	CreateAt  time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"` //软删除
// }
// type User2 struct {
// 	gorm.Model
// 	Name string
// }
// type Author struct {
// 	Name  string
// 	Email string
// }

// type Blog struct {
// 	ID      int
// 	Author  Author `gorm:"embedded"`
// 	Upvotes int32
// }

// func Model() {
// 	GL_DB.AutoMigrate(&Blog{}, &Author{})
// 	GL_DB.Create(&Blog{Author: Author{Name: "zhangsan123", Email: "123@163.com"}, Upvotes: 100})
// 	fmt.Println("插入数据成功")
// }
