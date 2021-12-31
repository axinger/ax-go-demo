package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	// 连接数据库
	open, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/ax_shop?characterEncoding=utf-8&useSSL=false")
	checkError(err)
	//插入数据
	//add(open)
	// 更新数据
	//update(open)
	// 删除数据
	del(open)
}

//插入数据
func add(open *sql.DB) {
	//插入数据
	prepare, err := open.Prepare("insert user set username=?,password=?,mobile=?,createtime=?")
	checkError(err)
	exec, err := prepare.Exec("李四", "123456", "17600000000", time.Now().Unix())
	checkError(err)
	id, err := exec.LastInsertId()
	checkError(err)
	fmt.Printf("插入数据ID: %d \n", id)
}

// 更新
func update(open *sql.DB) {
	prepare, err := open.Prepare("update user set username=? where id =?")
	checkError(err)
	exec, err := prepare.Exec("王五", "18")
	checkError(err)
	rows, err := exec.RowsAffected()
	checkError(err)
	fmt.Printf("更新数据成功,影响条数 %d \n", rows)
}

// 删除数据
func del(open *sql.DB) {
	prepare, err := open.Prepare("delete from user  where id =?")
	checkError(err)
	exec, err := prepare.Exec("8")
	checkError(err)
	rows, err := exec.RowsAffected()
	checkError(err)
	fmt.Printf("删除数据成功,影响条数 %d \n", rows)
}

//检测错误
func checkError(err error) {
	if err != nil {
		panic("操作失败:" + err.Error())
	}
}
