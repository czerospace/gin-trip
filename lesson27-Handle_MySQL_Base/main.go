package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//_ "gorm.io/driver/mysql"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.137.132:3306)/gin_demo")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 尝试建立连接
	if err := db.Ping(); err != nil {
		fmt.Println("connect to database failed...")
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	fmt.Println("connect to database success")
}
