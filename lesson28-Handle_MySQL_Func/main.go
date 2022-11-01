package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 封装 database 链接

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

func main() {
	if err := initializeDatabase(); err != nil {
		panic(err)
	}
	fmt.Println("connect to database success")
}
