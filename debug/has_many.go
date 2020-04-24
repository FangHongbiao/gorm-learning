package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}
type User struct {
	gorm.Model
	CreditCards []CreditCard
}

func main() {
	// 连接mysql数据库
	db, err := gorm.Open("mysql", "root:123456@(39.106.208.186:3309)/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db = db.Debug()

	fmt.Println("############Belongs To#################")
	fmt.Println("---------------建表看外键情况--------------")
	// struct和表对应起来
	db.AutoMigrate(&CreditCard{})
	db.AutoMigrate(&User{})
}
