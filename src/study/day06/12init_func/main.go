package main

import (
	// 匿名导入,需要 执行 init方法
	//_ "../10包"
	"fmt"
)

func init() {
	fmt.Println("12init初始化函数,init 方法,导入包自动执行")
}

func main() {
	fmt.Println("12init初始化函数 main")
}
