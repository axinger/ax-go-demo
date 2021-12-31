package main

import (
	"fmt"
)

// 查询数据 - 1条
func queryOne(id int) {

	sql := "SELECT * FROM t_student WHERE id=?;"
	var user userInfo
	// 含有关闭数据库方法,必须调用
	// 必须保证顺序
	//db.QueryRow(sql, id).
	//	Scan(&user.id,
	//		&user.name,
	//		&user.age,
	//		&user.sex,
	//		&user.address)

	// 预处理
	prepare, err := db.Prepare(sql)
	checkError(err)
	defer prepare.Close()

	// 执行参数
	prepare.QueryRow(id).
		Scan(&user.id,
			&user.name,
			&user.age,
			&user.sex,
			&user.address)

	fmt.Println("查询数据user:", user)

	fmt.Printf("查询数据user:%#v \n", user)

}

func queryMulti() {

	sql := "SELECT * FROM t_student"
	row, er := db.Query(sql)
	if er != nil {
		return
	}
	// ⚠️ 非常重要,关闭row,释放链接
	defer row.Close()

	var users []userInfo

	for row.Next() {

		var user userInfo
		// 含有关闭数据库方法,必须调用
		er := row.Scan(&user.id,
			&user.name,
			&user.age,
			&user.sex,
			&user.address)
		if er != nil {
			return
		}
		users = append(users, user)
	}
	fmt.Printf("查询数据user:%v \n", users)
	fmt.Printf("查询数据user:%#v \n", users)

}
