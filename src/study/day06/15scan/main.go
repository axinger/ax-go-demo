package main

import (
	"bufio"
	"fmt"
	"os"
)

func userScan() {

	var s string
	fmt.Print("请输入内容")
	fmt.Scan(&s) // 缺点, 不能识别空白 a b c 只能读取到a
	fmt.Printf("输入的内容是:%s\n", s)
}

func userBufio() {

	fmt.Print("请输入内容")

	/// 控制台输入
	reader := bufio.NewReader(os.Stdin)
	readString, _ := reader.ReadString('\n')

	fmt.Printf("输入的内容是:%s\n", readString)
}

func main() {
	//userScan()

	userBufio()
}
