package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var db *sqlx.DB

func initDB() error {

	var err error
	// 连接数据库
	db, err = sqlx.Connect("mysql", "root:12345678@tcp(127.0.0.1:3306)/ax_shop")

	if err != nil {
		log.Fatalln(err)
		fmt.Println("数据库链接失败======", err.Error())
		return err
	}

	db.SetMaxOpenConns(100) // 最大链接数
	//db.SetMaxIdleConns(5)   //最大空闲数量
	if err != nil {
		fmt.Println("数据库链接失败", err.Error())
		return err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("数据库链接验证密码失败", err.Error())
		return err
	}
	fmt.Println("数据库链接验证密码成功")

	return nil
}
