// package main

// import (
// 	"fmt"
// 	"math/rand"
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

// 	generateData(db, 1000)
// }

// func generateData(db *gorm.DB, count int) {

// 	for i := 0; i < count; i++ {

// 		name := fmt.Sprintf("user%d", rand.Intn(100))
// 		age := int64(10 + rand.Intn(50))
// 		timeStr := time.Now().Format("2006-01-02")
// 		today, _ := time.Parse("2006-01-02", timeStr)

// 		year := -5 + rand.Intn(5)
// 		month := 1 + rand.Intn(10)
// 		day := 1 + rand.Intn(25)
// 		birthday := today.AddDate(year, month, day)
// 		user := User{
// 			Name:     name,
// 			Age:      age,
// 			Birthday: birthday,
// 		}
// 		db.Create(&user)
// 	}
// }
