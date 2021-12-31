package main

import (
	"fmt"
)

func main() {

	err := initDB()
	if err != nil {
		fmt.Println("初始化数据库错误", err.Error())
	}
	checkError(err)

	queryOne(1)
	//queryMulti()

	/*	insert(userInfo{
		name:    "jack",
		age:     14,
		sex:     1,
		address: "街道",
	})*/

	/*	update(userInfo{
		id:      1,
		name:    "jack",
		age:     20,
		sex:     1,
		address: "街道",
	})*/

	//delete(2)

	updateTx(
		userInfo{
			id:  1,
			age: 24,
		},
		userInfo{
			id:  2,
			age: 22,
		},
	)
	//checkError(errors.New("测试错误"))
}

//检测错误
func checkError(err error) {
	if err != nil {
		panic("操作失败:" + err.Error())
		fmt.Println("操作失败:" + err.Error())
		return
	}
}
