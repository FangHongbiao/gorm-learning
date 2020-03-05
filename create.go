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
// 	Birthday time.Time
// }

// type Animal struct {
// 	ID   int64
// 	Name string `gorm:"default:'galeone'"`
// 	Age  int64
// }

// func (animal *Animal) BeforeCreate(scope *gorm.Scope) error {
// 	scope.SetColumn("name", "hahah")
// 	return nil
// }

// func main() {
// 	// 连接mysql数据库
// 	db, err := gorm.Open("mysql", "root:123456@(39.106.208.186:3306)/gorm?charset=utf8&parseTime=True&loc=Local")
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	defer db.Close()

// 	// struct和表对应起来
// 	db.AutoMigrate(&User{})
// 	db.AutoMigrate(&Animal{})

// 	// -------------创建记录------------------
// 	user := User{Name: "u2", Age: 18, Birthday: time.Now()}

// 	isNew := db.NewRecord(user) // => 主键为空返回`true`
// 	fmt.Println(isNew)

// 	db.Create(&user)

// 	isNew = db.NewRecord(user) // => 创建`user`后返回`false`
// 	fmt.Println(isNew)

// 	// -------------通过 tag 定义字段的默认值------------------
// 	// var animal = Animal{ID: 4, Age: 99, Name: ""}
// 	// db.Create(&animal)
// 	// fmt.Printf("%#v", animal)

// 	// ---------------在Hooks中设置字段值--------------------------
// 	// 使用了`func (animal *Animal) BeforeCreate(scope *gorm.Scope) error`
// 	var a1 = Animal{Age: 99, Name: ""}
// 	db.Create(&a1)
// 	fmt.Printf("%#v", a1)

// 	// ---------------为Instert语句添加扩展SQL选项---------------
// }
