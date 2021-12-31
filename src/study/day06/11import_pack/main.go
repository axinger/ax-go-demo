package main

import (
	// 取别名
	calcA "../10包"
	"fmt"
)

func main() {

	//add := Add(1, 2)
	add := calcA.Add(1, 2)
	fmt.Println("add = ", add)
}
