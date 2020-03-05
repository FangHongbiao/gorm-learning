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
// 	var users []User

// 	fmt.Println("#############################查询#############################")

// 	// 根据主键查询第一条记录
// 	user = User{}
// 	db.First(&user)
// 	fmt.Printf("根据主键查询第一条记录: %#v\n", user)
// 	//// SELECT * FROM users ORDER BY id LIMIT 1;

// 	// 随机获取一条记录
// 	user = User{}
// 	db.Take(&user)
// 	fmt.Printf("随机获取一条记录: %#v\n", user)
// 	//// SELECT * FROM users LIMIT 1;

// 	// 根据主键查询最后一条记录
// 	user = User{}
// 	db.Last(&user)
// 	fmt.Printf("根据主键查询最后一条记录: %#v\n", user)
// 	//// SELECT * FROM `users`  WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = 1 ORDER BY `users`.`id` DESC LIMIT 1;

// 	// 查询所有的记录
// 	users = []User{}
// 	db.Find(&users)
// 	fmt.Printf("查询所有的记录: %#v\n", users)
// 	//// SELECT * FROM users;

// 	// 查询指定的某条记录(仅当主键为整型时可用)
// 	user = User{}
// 	db.First(&user, 10)
// 	fmt.Printf("查询指定的某条记录(仅当主键为整型时可用): %#v\n", user)
// 	//// SELECT * FROM users WHERE id = 10;

// 	fmt.Println("---------------Where 条件-----------------")

// 	// 1. 普通 SQL

// 	// Get first matched record
// 	user = User{}
// 	db.Where("name = ?", "user50").First(&user)
// 	//// SELECT * FROM users WHERE name = 'user50' limit 1;

// 	// Get all matched records
// 	users = []User{}
// 	db.Where("name = ?", "user50").Find(&users)
// 	//// SELECT * FROM users WHERE name = 'user50';

// 	// <>
// 	users = []User{}
// 	db.Where("name <> ?", "user50").Find(&users)
// 	//// SELECT * FROM users WHERE name <> 'user50';

// 	// IN
// 	users = []User{}
// 	db.Where("name IN (?)", []string{"user50", "user51", "user52"}).Find(&users)
// 	//// SELECT * FROM users WHERE name in ("user50", "user51", "user52");

// 	// LIKE
// 	users = []User{}
// 	db.Where("name LIKE ?", "%55%").Find(&users)
// 	//// SELECT * FROM users WHERE name LIKE '%55%';

// 	// AND
// 	users = []User{}
// 	db.Where("name = ? AND age >= ?", "user50", "22").Find(&users)
// 	//// SELECT * FROM users WHERE name = 'user50' AND age >= 22;

// 	// Time
// 	timeStr := time.Now().Format("2006-01-02")
// 	today, _ := time.Parse("2006-01-02", timeStr)
// 	lastWeek := today.AddDate(0, 0, -7)
// 	fmt.Println(today, lastWeek)

// 	users = []User{}
// 	db.Where("updated_at > ?", lastWeek).Find(&users)
// 	//// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

// 	// BETWEEN
// 	users = []User{}
// 	db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
// 	//// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

// 	// 2. Struct & Map

// 	// Struct
// 	user = User{}
// 	db.Where(&User{Name: "user88", Age: 20}).First(&user)
// 	//// SELECT * FROM users WHERE name = "user88" AND age = 20 LIMIT 1;

// 	// Map
// 	users = []User{}
// 	db.Where(map[string]interface{}{"name": "user88", "age": 20}).Find(&users)
// 	//// SELECT * FROM users WHERE name = "user88" AND age = 20;

// 	// 主键的切片
// 	users = []User{}
// 	db.Where([]int64{20, 21, 22}).Find(&users)
// 	//// SELECT * FROM users WHERE id IN (20, 21, 22);

// 	fmt.Println("--------------------Not 条件---------------------")

// 	users = []User{}
// 	db.Not("name", "user88").First(&user)
// 	//// SELECT * FROM users WHERE name <> "user88" LIMIT 1;

// 	// Not In
// 	users = []User{}
// 	db.Not("name", []string{"user87", "user88"}).Find(&users)
// 	//// SELECT * FROM users WHERE name NOT IN ("user88", "user87");

// 	// Not In slice of primary keys
// 	user = User{}
// 	db.Not([]int64{1, 2, 3}).First(&user)
// 	//// SELECT * FROM users WHERE id NOT IN (1,2,3);

// 	user = User{}
// 	db.Not([]int64{}).First(&user)
// 	//// SELECT * FROM users;

// 	// Plain SQL
// 	user = User{}
// 	db.Not("name = ?", "user88").First(&user)
// 	//// SELECT * FROM users WHERE NOT(name = "jinzhu");

// 	// Struct
// 	user = User{}
// 	db.Not(User{Name: "user88"}).First(&user)
// 	//// SELECT * FROM users WHERE name <> "jinzhu";

// 	fmt.Println("----------------Or 条件------------------------")

// 	users = []User{}
// 	db.Where("name = ?", "user88").Or("name = ?", "user89").Find(&users)
// 	//// SELECT * FROM users WHERE name = 'user88' OR name = 'user89';

// 	// Struct
// 	users = []User{}
// 	db.Where("name = 'user88'").Or(User{Name: "user89"}).Find(&users)
// 	//// SELECT * FROM users WHERE name = 'user88' OR name = 'user89';

// 	// Map
// 	users = []User{}
// 	db.Where("name = 'user89'").Or(map[string]interface{}{"name": "user88"}).Find(&users)
// 	//// SELECT * FROM users WHERE name = 'jinzhu' OR name = 'jinzhu 2';

// 	fmt.Println("--------------Inline Condition 内联条件-------------------")

// 	// 根据主键获取记录 (只适用于整形主键)
// 	db.First(&user, 23)
// 	//// SELECT * FROM users WHERE id = 23 LIMIT 1;

// 	// 根据主键获取记录, 如果它是一个非整形主键
// 	user = User{}
// 	db.First(&user, "id = ?", "44")
// 	//// SELECT * FROM users WHERE id = 'string_primary_key' LIMIT 1;

// 	// Plain SQL
// 	user = User{}
// 	db.Find(&user, "name = ?", "user88")
// 	//// SELECT * FROM users WHERE name = "jinzhu";

// 	users = []User{}
// 	db.Find(&users, "name <> ? AND age > ?", "user88", 20)
// 	//// SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;

// 	// Struct
// 	users = []User{}
// 	db.Find(&users, User{Age: 20})
// 	//// SELECT * FROM users WHERE age = 20;

// 	// Map
// 	users = []User{}
// 	db.Find(&users, map[string]interface{}{"age": 20})
// 	//// SELECT * FROM users WHERE age = 20;

// 	fmt.Println("------------Extra Querying option 其它查询选项-----------")

// 	// 为查询 SQL 添加额外的 SQL 操作
// 	user = User{}
// 	db.Set("gorm:query_option", "FOR UPDATE").First(&user, 10)
// 	//// SELECT * FROM users WHERE id = 10 FOR UPDATE;

// 	fmt.Println("################FirstOrInit#####################")

// 	// 未找到
// 	user = User{}
// 	db.FirstOrInit(&user, User{Name: "non_existing"})
// 	//// user -> User{Name: "non_existing"}

// 	// 找到
// 	user = User{}
// 	db.Where(User{Name: "user88"}).FirstOrInit(&user)
// 	//// user -> User{Id: 111, Name: "user88", Age: 20}
// 	user = User{}
// 	db.FirstOrInit(&user, map[string]interface{}{"name": "user88"})
// 	//// user -> User{Id: 111, Name: "user88", Age: 20}

// 	fmt.Println("---------------Attrs----------------")

// 	// 未找到
// 	user = User{}
// 	db.Where(User{Name: "non_existing"}).Attrs(User{Age: 20}).FirstOrInit(&user)
// 	//// SELECT * FROM USERS WHERE name = 'non_existing';
// 	//// user -> User{Name: "non_existing", Age: 20}

// 	user = User{}
// 	db.Where(User{Name: "non_existing"}).Attrs("age", 20).FirstOrInit(&user)
// 	//// SELECT * FROM USERS WHERE name = 'non_existing';
// 	//// user -> User{Name: "non_existing", Age: 20}

// 	// 找到
// 	user = User{}
// 	db.Where(User{Name: "user88"}).Attrs(User{Age: 30}).FirstOrInit(&user)
// 	//// SELECT * FROM USERS WHERE name = user88';
// 	//// user -> User{Id: 111, Name: "Jinzhu", Age: 20}

// 	fmt.Println("---------------Assign---------------------")

// 	// 未找到
// 	user = User{}
// 	// 未找到
// 	db.Where(User{Name: "non_existing"}).Assign(User{Age: 20}).FirstOrInit(&user)
// 	//// user -> User{Name: "non_existing", Age: 20}

// 	// 找到
// 	user = User{}
// 	db.Where(User{Name: "Jinzhu"}).Assign(User{Age: 30}).FirstOrInit(&user)
// 	//// SELECT * FROM USERS WHERE name = jinzhu';
// 	//// user -> User{Id: 111, Name: "Jinzhu", Age: 30}

// 	fmt.Println("###############FirstOrCreate#################")

// 	// 未找到
// 	user = User{}
// 	db.FirstOrCreate(&user, User{Name: "non_existing", Birthday: time.Now()})
// 	//// INSERT INTO "users" (name) VALUES ("non_existing");
// 	//// user -> User{Id: 112, Name: "non_existing"}

// 	// 找到
// 	user = User{}
// 	birth, _ := time.Parse("2006-01-02 15:04:05", "2016-11-18 08:00:00")
// 	db.Where(User{Name: "user50", Birthday: birth}).FirstOrCreate(&user)
// 	//// user -> User{Id: 111, Name: "user50"}

// 	fmt.Println("------------------Attrs-------------------")

// 	// 未找到
// 	user = User{}
// 	db.Where(User{Name: "non_existing1"}).Attrs(User{Age: 20}).FirstOrCreate(&user)
// 	//// SELECT * FROM users WHERE name = 'non_existing';
// 	//// INSERT INTO "users" (name, age) VALUES ("non_existing", 20);
// 	//// user -> User{Id: 112, Name: "non_existing", Age: 20}

// 	// 找到
// 	user = User{}
// 	db.Where(User{Name: "jinzhu"}).Attrs(User{Age: 30}).FirstOrCreate(&user)
// 	//// SELECT * FROM users WHERE name = 'jinzhu';
// 	//// user -> User{Id: 111, Name: "jinzhu", Age: 20}

// 	fmt.Println("----------------Assign----------------------")

// 	// 未找到
// 	user = User{}
// 	db.Where(User{Name: "non_existing2"}).Assign(User{Age: 20}).FirstOrCreate(&user)
// 	//// SELECT * FROM users WHERE name = 'non_existing';
// 	//// INSERT INTO "users" (name, age) VALUES ("non_existing", 20);
// 	//// user -> User{Id: 112, Name: "non_existing", Age: 20}

// 	// 找到
// 	user = User{}
// 	db.Where(User{Name: "jinzhu"}).Assign(User{Age: 30}).FirstOrCreate(&user)
// 	//// SELECT * FROM users WHERE name = 'jinzhu';
// 	//// UPDATE users SET age=30 WHERE id = 111;
// 	//// user -> User{Id: 111, Name: "jinzhu", Age: 30}

// 	fmt.Println("##############Advanced Query 高级查询##################")
// 	fmt.Println("----------SubQuery 子查询----------------")
// 	// db.Where("amount > ?", DB.Table("orders").Select("AVG(amount)").Where("state = ?", "paid").QueryExpr()).Find(&orders)
// 	// // SELECT * FROM "orders"  WHERE "orders"."deleted_at" IS NULL AND (amount > (SELECT AVG(amount) FROM "orders"  WHERE (state = 'paid')));

// 	fmt.Println("------------选择字段---------------")
// 	users = []User{}
// 	db.Select("name, age").Find(&users)
// 	//// SELECT name, age FROM users;

// 	users = []User{}
// 	db.Select([]string{"name", "age"}).Find(&users)
// 	//// SELECT name, age FROM users;

// 	db.Table("users").Select("COALESCE(age,?)", 42).Rows()
// 	//// SELECT COALESCE(age,'42') FROM users;

// 	fmt.Println("-------------排序---------------")
// 	users = []User{}
// 	db.Order("age desc, name").Find(&users)
// 	//// SELECT * FROM users ORDER BY age desc, name;

// 	// 多字段排序
// 	users = []User{}
// 	db.Order("age desc").Order("name").Find(&users)
// 	//// SELECT * FROM users ORDER BY age desc, name;

// 	// 覆盖排序
// 	users1 := []User{}
// 	users2 := []User{}
// 	db.Order("age desc").Find(&users1).Order("age", true).Find(&users2)
// 	//// SELECT * FROM users ORDER BY age desc; (users1)
// 	//// SELECT * FROM users ORDER BY age; (users2)

// 	fmt.Println("--------------数量-------------------")
// 	users = []User{}
// 	db.Limit(3).Find(&users)
// 	//// SELECT * FROM users LIMIT 3;

// 	// -1 取消 Limit 条件
// 	users1 = []User{}
// 	users2 = []User{}
// 	db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
// 	//// SELECT * FROM users LIMIT 10; (users1)
// 	//// SELECT * FROM users; (users2)

// 	// TODO 偏移(MYSQL似乎无效)
// 	fmt.Println("--------------偏移(MYSQL似乎无效)-------------------")
// 	users = []User{}
// 	db.Offset(3).Find(&users)
// 	//// SELECT * FROM users OFFSET 3;

// 	// -1 取消 Offset 条件
// 	users1 = []User{}
// 	users2 = []User{}
// 	db.Offset(10).Find(&users1).Offset(-1).Find(&users2)
// 	//// SELECT * FROM users OFFSET 10; (users1)
// 	//// SELECT * FROM users; (users2)

// 	fmt.Println("------------总数-------------------")
// 	users = []User{}
// 	var count int
// 	db.Where("name = ?", "user82").Or("name = ?", "user80").Find(&users).Count(&count)
// 	fmt.Println(count)
// 	//// SELECT * from USERS WHERE name = 'jinzhu' OR name = 'jinzhu 2'; (users)
// 	//// SELECT count(*) FROM users WHERE name = 'jinzhu' OR name = 'jinzhu 2'; (count)

// 	db.Model(&User{}).Where("name = ?", "user80").Count(&count)
// 	fmt.Println(count)
// 	//// SELECT count(*) FROM users WHERE name = 'user80'; (count)

// 	db.Table("users").Count(&count)
// 	fmt.Println(count)
// 	//// SELECT count(*) FROM users;

// 	db.Table("users").Select("count(distinct(name))").Count(&count)
// 	fmt.Println(count)
// 	//// SELECT count( distinct(name) ) FROM users; (count)

// 	fmt.Println("---------------Group & Having--------------")
// 	// rows, err := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Rows()
// 	// for rows.Next() {
// 	// 	...
// 	// }

// 	// rows, err := db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Rows()
// 	// for rows.Next() {
// 	// 	...
// 	// }

// 	// type Result struct {
// 	// 	Date  time.Time
// 	// 	Total int64
// 	// }
// 	// db.Table("orders").Select("date(created_at) as date, sum(amount) as total").Group("date(created_at)").Having("sum(amount) > ?", 100).Scan(&results)

// 	fmt.Println("--------------连接------------")
// 	// rows, err := db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Rows()
// 	// for rows.Next() {
// 	// 	...
// 	// }

// 	// db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)

// 	// // 多连接及参数
// 	// db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("credit_cards.number = ?", "411111111111").Find(&user)

// 	fmt.Println("###Pluck: Pluck，查询 model 中的一个列作为切片，如果您想要查询多个列，您应该使用 Scan###")
// 	var ages []int64
// 	db.Find(&users).Pluck("age", &ages)

// 	// TODO 注意下面量化总写法的不同
// 	var names []string
// 	db.Model(&User{}).Pluck("name", &names)
// 	// SELECT name FROM `users`  WHERE `users`.`deleted_at` IS NULL

// 	db.Table("users").Pluck("name", &names)
// 	// SELECT name FROM `users

// 	// 想查询多个字段？ 这样做：
// 	db.Select("name, age").Find(&users)

// 	fmt.Println("-------------扫描----------------")
// 	type Result struct {
// 		Name string
// 		Age  int64
// 	}

// 	var result Result
// 	db.Table("users").Select("name, age").Where("name = ?", "user12").Scan(&result)

// 	// 原生 SQL
// 	db.Raw("SELECT name, age FROM users WHERE name = ?", "user12").Scan(&result)
// }
