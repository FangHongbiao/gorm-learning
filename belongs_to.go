// package main

// import (
// 	"fmt"

// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// )

// // belongs_to
// type User struct {
// 	gorm.Model
// 	Name string
// }

// // `Profile` belongs to `User`, `UserID` is the foreign key
// type Profile struct {
// 	gorm.Model
// 	UserID int
// 	User   User
// 	Name   string
// }

// // 外键
// // type User struct {
// // 	gorm.Model
// // 	Name string
// // }

// // type Profile struct {
// // 	gorm.Model
// // 	Name      string
// // 	User      User `gorm:"foreignkey:UserRefer"` // use UserRefer as foreign key
// // 	UserRefer uint
// // }

// // 关联外键
// // type User struct {
// // 	gorm.Model
// // 	Refer string
// // 	Name  string
// // }

// // type Profile struct {
// // 	gorm.Model
// // 	Name      string
// // 	User      User `gorm:"association_foreignkey:Refer"` // 将 Refer 作为关联外键
// // 	UserRefer string
// // }

// // Has one
// // User 只能有一张信用卡 (CreditCard), CreditCardID 是外键

// func main() {
// 	// 连接mysql数据库
// 	db, err := gorm.Open("mysql", "root:123456@(39.106.208.186:3306)/gorm?charset=utf8&parseTime=True&loc=Local")
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	defer db.Close()

// 	db = db.Debug()

// 	fmt.Println("############Belongs To#################")
// 	fmt.Println("---------------建表看外键情况--------------")
// 	// struct和表对应起来
// 	db.AutoMigrate(&User{})
// 	db.AutoMigrate(&Profile{})

// 	fmt.Println(("-------------------Belongs To 的使用-------------------"))
// 	user := User{}
// 	user.ID = 111
// 	profile := Profile{}
// 	db.Model(&user).Related(&profile)
// 	//// SELECT * FROM profiles WHERE user_id = 111; // 111 is user's ID

// }
