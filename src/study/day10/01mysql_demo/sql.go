package main

import (
	"database/sql"
)

// 增删改
func Excud(sql string, args ...interface{}) (sql.Result, error) {

	prepare, err := db.Prepare(sql)
	defer prepare.Close()

	result, err := prepare.Exec(args)
	return result, err
}
