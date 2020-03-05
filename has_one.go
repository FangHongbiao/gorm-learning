// package main

// import (
// 	"fmt"

// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// )

// // 默认外键 UserID
// type CreditCard struct {
// 	gorm.Model
// 	Number string
// 	UserID uint
// }

// type User struct {
// 	gorm.Model
// 	CreditCard CreditCard
// }

// // // 指定外键
// // type CreditCard struct {
// // 	gorm.Model
// // 	Number   string
// // 	UserName string
// // }

// // type User struct {
// // 	gorm.Model
// // 	CreditCard CreditCard `gorm:"foreignkey:UserName"`
// // }

// // // 关联外键
// // type CreditCard struct {
// // 	gorm.Model
// // 	Number string
// // 	UID    string
// // }

// // type User struct {
// // 	gorm.Model
// // 	Name       `sql:"index"`
// // 	CreditCard CreditCard `gorm:"foreignkey:uid;association_foreignkey:name"`
// // }

// // 多态关联
// type Cat struct {
// 	ID   int
// 	Name string
// 	Toy  Toy `gorm:"polymorphic:Owner;"`
// }

// type Dog struct {
// 	ID   int
// 	Name string
// 	Toy  Toy `gorm:"polymorphic:Owner;"`
// }

// type Toy struct {
// 	ID        int
// 	Name      string
// 	OwnerID   int
// 	OwnerType string
// }

// func main() {
// 	// 连接mysql数据库
// 	db, err := gorm.Open("mysql", "root:123456@(39.106.208.186:3309)/gorm?charset=utf8&parseTime=True&loc=Local")
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	defer db.Close()

// 	db = db.Debug()

// 	fmt.Println("############Has One#################")
// 	fmt.Println("---------------建表看外键情况--------------")

// 	db.AutoMigrate(&CreditCard{})
// 	db.AutoMigrate(&User{})

// 	// 多态关联
// 	db.AutoMigrate(&Dog{})
// 	db.AutoMigrate(&Cat{})
// 	db.AutoMigrate(&Toy{})

// 	fmt.Println("--------------Has One 的使用---------------")
// 	var card CreditCard
// 	user := User{}
// 	user.ID = 123
// 	db.Model(&user).Related(&card, "CreditCard")
// 	//// SELECT * FROM credit_cards WHERE user_id = 123; // 123 is user's primary key
// 	// CreditCard 是 users 的字段，其含义是，获取 user 的 CreditCard 并填充至 card 变量
// 	// 如果字段名与 model 名相同，比如上面的例子，此时字段名可以省略不写，像这样：
// 	db.Model(&user).Related(&card)
// }
