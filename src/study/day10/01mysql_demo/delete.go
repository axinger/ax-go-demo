package main

import (
	"fmt"
)

// 删除数据
func delete(id int) {

	/*

		go mysql prepare_Go 语言操作 MySQL 之 预处理

			与以往的mysql操作相似，也是一个连接对象，对于插入，删除等要先prepare，然后执行Exec，在exec阶段绑定占位符变量。

			预处理优点

			预处理语句大大减少了分析时间，只做了一次查询(虽然语句多次执行)；

			绑定参数减少了服务器带宽，只需发送查询的参数，而不是整个语句；

			预处理语句针对 SQL 注入是非常有用的，因为参数值发送后使用不同的协议，保证了数据的合法性。
	*/

	sql := "delete from t_student  where id =?"
	prepare, err := db.Prepare(sql)
	checkError(err)
	exec, err := prepare.Exec(id)
	checkError(err)
	rows, err := exec.RowsAffected()
	checkError(err)
	fmt.Printf("删除数据成功,影响条数 %d \n", rows)

}
