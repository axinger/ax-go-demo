package main

import (
	"fmt"
)

// 更新
func update(user userInfo) {

	sql := "update t_student set age=? where id =?"
	//result, err := db.Exec(sql, user.age, user.id)
	//if err != nil {
	//	return
	//}

	prepare, err := db.Prepare(sql)
	checkError(err)
	defer prepare.Close()

	result, err := prepare.Exec(user.age, user.id)
	// 更新数据,操作的行数
	rows, err := result.RowsAffected()
	if err != nil {
		return
	}
	fmt.Printf("更新数据成功,影响条数 %d \n", rows)
}

//更新数据
func updateTx(user1 userInfo, user2 userInfo) {

	// 事务
	tx, err := db.Begin()
	if err != nil {
		return
	}

	{

		sql := "update t_student set age=? where id =?"

		// 预处理,针对 sql 相同的,事务的时候,可能不同表操作,sql不同
		prepare, err := tx.Prepare(sql)
		if err != nil {
			fmt.Println("更新数据的失败1,需要回滚2:", err.Error())
			tx.Rollback()
			return
		}
		// 关闭链接
		defer prepare.Close()

		// 预处理后,执行语句 直接放入参数,不要再放sql
		result, err := prepare.Exec(user1.age, user1.id)
		if err != nil {
			fmt.Println("更新数据的失败1,需要回滚1:", err.Error())
			tx.Rollback()
			return
		}

		// 如果是更新数据的操作,可以得到更新数据的id
		id, err := result.LastInsertId()
		if err != nil {
			fmt.Println("更新数据的失败2:", err.Error())
			return
		}
		fmt.Printf("更新数据的ID: %d \n", id)

	}

	{

		sql := "update xxx set age=? where id =?"

		// 预处理,针对 sql 相同的,事务的时候,可能不同表操作,sql不同
		prepare, err := tx.Prepare(sql)

		if err != nil {
			fmt.Println("更新数据的失败1,需要回滚2:", err.Error())
			tx.Rollback()
			return
		}

		// 关闭链接
		defer prepare.Close()

		// 预处理后,执行语句 直接放入参数,不要再放sql
		result, err := prepare.Exec(user2.age, user2.id)
		if err != nil {
			fmt.Println("更新数据的失败1,需要回滚2:", err.Error())
			tx.Rollback()
			return
		}

		// 如果是更新数据的操作,可以得到更新数据的id
		id, err := result.LastInsertId()
		if err != nil {
			fmt.Println("更新数据的失败2:", err.Error())
			return
		}
		fmt.Printf("更新数据的ID: %d \n", id)

	}

	// 事务提交
	tx.Commit()
}
