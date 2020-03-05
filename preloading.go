package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// // 下面的例子会用到 User 和 Order 结构体
// type User struct {
// 	gorm.Model
// 	Username string
// 	Orders   Order
// }
// type Order struct {
// 	gorm.Model
// 	UserID uint
// 	Price  float64
// 	State  int
// }

func main() {
	// 连接mysql数据库
	db, err := gorm.Open("mysql", "root:123456@(39.106.208.186:3306)/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db = db.Debug()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Order{})

	var users []User

	fmt.Println("############Belongs To#################")

	fmt.Println("-----------------预加载----------------------")
	// Preload 方法的参数应该是主结构体的字段名
	users = []User{}
	db.Preload("Orders").Find(&users)
	//// SELECT * FROM users;
	//// SELECT * FROM orders WHERE user_id IN (1,2,3,4);

	users = []User{}
	db.Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
	//// SELECT * FROM users;
	//// SELECT * FROM orders WHERE user_id IN (1,2,3,4) AND state NOT IN ('cancelled');

	users = []User{}
	db.Where("state = ?", "active").Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
	//// SELECT * FROM users WHERE state = 'active';
	//// SELECT * FROM orders WHERE user_id IN (1,2) AND state NOT IN ('cancelled');

	// users = []User{}
	// db.Preload("Orders").Preload("Profile").Preload("Role").Find(&users)
	//// SELECT * FROM users;
	//// SELECT * FROM orders WHERE user_id IN (1,2,3,4); // has many
	//// SELECT * FROM profiles WHERE user_id IN (1,2,3,4); // has one
	//// SELECT * FROM roles WHERE id IN (4,5,6); // belongs to

	fmt.Println("---------------自动预加载----------------")
	db.Set("gorm:auto_preload", true).Find(&users)
}
