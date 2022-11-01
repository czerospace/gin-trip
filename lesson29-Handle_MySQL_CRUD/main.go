package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// CRUD

var db *sql.DB

func initializeDatabase() (err error) {
	dsn := "root:123456@tcp(192.168.137.132:3306)/gin_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return
	}
	return nil
}

// 定义一个结构体
type user struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Age  uint32 `json:"age"`
}

// 查询单条数据
func querySingleRow() user {
	sqlStr := "select * from student where id = ?"
	var u user
	if err := db.QueryRow(sqlStr, 1).Scan(&u.Id, &u.Name, &u.Age); err != nil {
		log.Printf("scan failed err: %v\n", err)
		return u
	}
	log.Println(u.Id, u.Name, u.Age)
	return u
}

// 查询多条数据
func queryMultiRow() []user {
	sqlstr := "select * from student"
	rows, err := db.Query(sqlstr)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()
	// 创建一个 slice
	users := make([]user, 0)
	// 遍历多行数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.Id, &u.Name, &u.Age)
		if err != nil {
			log.Println(err)
			return nil
		}
		users = append(users, u)
	}
	return users
}

// 更新数据
func updateRow() {
	sqlStr := "update student set name = ? where id = ?"
	res, err := db.Exec(sqlStr, "niko", 1)
	if err != nil {
		fmt.Printf("update failed err: %v\n", err)
	}
	// 受影响的行数
	n, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed %v\n", err)
	}
	fmt.Println(n)
	fmt.Println("update success")
}

func main() {
	if err := initializeDatabase(); err != nil {
		panic(err)
	}
	fmt.Println("connect to database success")
	//r := gin.Default()
	//// 单个用户查询
	//r.GET("user", func(c *gin.Context) {
	//	u := querySingleRow()
	//	c.JSON(200, gin.H{
	//		"data": u,
	//	})
	//})
	//// 查询用户列表
	//r.GET("users", func(c *gin.Context) {
	//	u := queryMultiRow()
	//	c.JSON(200, gin.H{
	//		"data": u,
	//	})
	//})
	//err := r.Run()
	//if err != nil {
	//	return
	//}

	// 更新用户
	updateRow()
}
