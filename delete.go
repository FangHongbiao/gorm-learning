// package main

// import (
// 	"fmt"
// 	"time"

// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// )

// type User struct {
// 	gorm.Model
// 	Name     string
// 	Age      int64
// 	Actived  bool      `orm:"DEFAULT:true`
// 	Birthday time.Time `gorm:"DEFAULT:'2020-03-01 14:47:22'"`
// }

// func main() {
// 	// 连接mysql数据库
// 	db, err := gorm.Open("mysql", "root:123456@(39.106.208.186:3306)/gorm?charset=utf8&parseTime=True&loc=Local")
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	defer db.Close()
// 	// 开启调试，打印SQL语句
// 	db = db.Debug()
// 	// struct和表对应起来
// 	db.AutoMigrate(&User{})

// 	var user User
// 	// var users []User

// 	fmt.Println("############删除记录#############")
// 	// 删除现有记录
// 	user = User{}
// 	user.ID = 2
// 	db.Delete(&user)
// 	//// DELETE from users where id=10;

// 	// TODO MYSQL下语法错误
// 	// 为删除 SQL 添加额外的 SQL 操作
// 	// user = User{}
// 	// user.ID = 2
// 	// db.Set("gorm:delete_option", "OPTION (OPTIMIZE FOR UNKNOWN)").Delete(&user)
// 	//// DELETE from users where id=10 OPTION (OPTIMIZE FOR UNKNOWN);

// 	fmt.Println("-----------------批量删除-----------------")
// 	db.Where("name LIKE ?", "%hell%").Delete(User{})
// 	//// DELETE from users where name LIKE "%jinzhu%";

// 	db.Delete(User{}, "name LIKE ?", "%hell%")
// 	//// DELETE from users where name LIKE "%jinzhu%";

// 	fmt.Println("-----------------软删除-----------------")
// 	user = User{}
// 	user.ID = 111
// 	db.Delete(&user)
// 	//// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE id = 111;

// 	// 批量删除
// 	db.Where("age = ?", 20).Delete(&User{})
// 	//// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE age = 20;

// 	// 查询记录时会忽略被软删除的记录
// 	user = User{}
// 	db.Where("age = 20").Find(&user)
// 	//// SELECT * FROM users WHERE age = 20 AND deleted_at IS NULL;

// 	// Unscoped 方法可以查询被软删除的记录
// 	users := []User{}
// 	db.Unscoped().Where("age = 20").Find(&users)
// 	//// SELECT * FROM users WHERE age = 20;

// 	fmt.Println("-----------------物理删除-----------------")
// 	user = User{}
// 	user.ID = 1
// 	// Unscoped 方法可以物理删除记录
// 	db.Unscoped().Delete(&user)
// 	//// DELETE FROM orders WHERE id=10;
// }
