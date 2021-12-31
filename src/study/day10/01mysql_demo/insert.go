package main

import (
	"fmt"
)

//插入数据
func insert(user userInfo) {

	//插入数据
	sql := "insert t_student set name=?,age=?,sex=?,address=?"
	//
	//result, err := db.Exec(sql, user.name, user.age, user.sex, user.address)

	// 预处理
	prepare, err := db.Prepare(sql)
	checkError(err)
	// 关闭链接
	defer prepare.Close()

	// 预处理后,执行语句 直接放入参数,不要再放sql
	result, err := prepare.Exec(user.name, user.age, user.sex, user.address)

	if err != nil {
		fmt.Println("插入数据的失败1:", err.Error())
		return
	}
	// 如果是插入数据的操作,可以得到插入数据的id
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("插入数据的失败2:", err.Error())
		return
	}
	fmt.Printf("插入数据的ID: %d \n", id)
}
