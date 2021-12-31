package calc

import "fmt"

func Add(a, b int) int {

	return a + b
}

func init() {
	fmt.Println("10包的,init 方法,导入包自动执行,可以同时存在多个===calc.go")
}
