package main

import (
	"fmt"
)

// 查询数据 - 1条
func queryOne(id int) {
	sql := "SELECT * FROM t_student WHERE id=?;"
	var user userInfo
	err := db.Get(&user, sql, id)
	if err != nil {
		fmt.Println("查询一个数据错误:" + err.Error())
		return
	}
	fmt.Println("查询一个数据user:", user)
	fmt.Printf("查询一个数据user:%#v \n", user)

}

func queryMulti() {

	sql := "SELECT * FROM t_student"

	var users []userInfo

	err := db.Select(&users, sql)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("查询多个数据user:%v \n", users)
	fmt.Printf("查询多个数据user:%#v \n", users)

}
