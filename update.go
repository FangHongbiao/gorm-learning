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
// 	Actived  bool `orm:"DEFAULT:true`
// 	Birthday time.Time `gorm:"DEFAULT:'2020-03-01 14:47:22'"`
// }

// type Product struct {
// 	gorm.Model
// 	Price    int64
// 	Quantity int64
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
// 	db.AutoMigrate(&Product{})

// 	var user User
// 	// var users []User

// 	fmt.Println("########更新所有字段###########")
// 	db.First(&user)

// 	user.Name = "jinzhu 2"
// 	user.Age = 100
// 	db.Save(&user)
// 	//// UPDATE users SET name='jinzhu 2', age=100, birthday='2016-01-01', updated_at = '2013-11-17 21:34:10' WHERE id=111;

// 	fmt.Println("########更新修改字段###########")
// 	// 更新单个属性，如果它有变化
// 	db.Model(&user).Update("name", "hello")
// 	//// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

// 	// 根据给定的条件更新单个属性
// 	db.Model(&user).Where("active = ?", true).Update("name", "hello")
// 	//// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;

// 	// 使用 map 更新多个属性，只会更新其中有变化的属性
// 	db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
// 	//// UPDATE users SET name='hello', age=18, actived=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

// 	// 使用 struct 更新多个属性，只会更新其中有变化且为非零值的字段
// 	db.Model(&user).Updates(User{Name: "hello", Age: 18})
// 	//// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

// 	// 警告：当使用 struct 更新时，GORM只会更新那些非零值的字段
// 	// 对于下面的操作，不会发生任何更新，"", 0, false 都是其类型的零值
// 	db.Model(&user).Updates(User{Name: "", Age: 0, Actived: false})

// 	fmt.Println("########更新选定字段###########")
// 	db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
// 	//// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

// 	db.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "actived": false})
// 	//// UPDATE users SET age=18, actived=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

// 	fmt.Println("########无 Hooks 更新###########")
// 	// 更新单个属性，类似于 `Update`
// 	db.Model(&user).UpdateColumn("name", "hello")
// 	//// UPDATE users SET name='hello' WHERE id = 111;

// 	// 更新多个属性，类似于 `Updates`
// 	db.Model(&user).UpdateColumns(User{Name: "hello", Age: 18})
// 	//// UPDATE users SET name='hello', age=18 WHERE id = 111;

// 	fmt.Println("########批量更新###########")
// 	db.Table("users").Where("id IN (?)", []int{10, 11}).Updates(map[string]interface{}{"name": "hello", "age": 18})
// 	//// UPDATE users SET name='hello', age=18 WHERE id IN (10, 11);

// 	// 使用 struct 更新时，只会更新非零值字段，若想更新所有字段，请使用map[string]interface{}
// 	db.Model(User{}).Updates(User{Name: "hello", Age: 18})
// 	//// UPDATE users SET name='hello', age=18;

// 	// 使用 `RowsAffected` 获取更新记录总数
// 	// db.Model(User{}).Updates(User{Name: "hello", Age: 18}).RowsAffected

// 	fmt.Println("########使用 SQL 表达式更新###########")
// 	var product = Product{Price: 10, Quantity: 500}

// 	db.Create(&product)

// 	db.Model(&product).Update("price", gorm.Expr("price * ? + ?", 2, 100))
// 	//// UPDATE "products" SET "price" = price * '2' + '100', "updated_at" = '2013-11-17 21:34:10' WHERE "id" = '2';

// 	db.Model(&product).Updates(map[string]interface{}{"price": gorm.Expr("price * ? + ?", 2, 100)})
// 	//// UPDATE "products" SET "price" = price * '2' + '100', "updated_at" = '2013-11-17 21:34:10' WHERE "id" = '2';

// 	db.Model(&product).UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
// 	//// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = '2';

// 	db.Model(&product).Where("quantity > 1").UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
// 	//// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = '2' AND quantity > 1;

// 	fmt.Println("########修改 Hooks 中的值###########")
// 	// func (user *User) BeforeSave(scope *gorm.Scope) (err error) {
// 	// 	if pw, err := bcrypt.GenerateFromPassword(user.Password, 0); err == nil {
// 	// 		scope.SetColumn("EncryptedPassword", pw)
// 	// 	}
// 	// }

// 	fmt.Println("########其它更新选项###########")
// 	// // 为 update SQL 添加其它的 SQL
// 	// db.Model(&user).Set("gorm:update_option", "OPTION (OPTIMIZE FOR UNKNOWN)").Update("name", "hello")
// 	// //// UPDATE users SET name='hello', updated_at = '2013-11-17 21:34:10' WHERE id=111 OPTION (OPTIMIZE FOR UNKNOWN);

// }
