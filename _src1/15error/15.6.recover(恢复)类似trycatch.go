package main

import "fmt"

/*

Go语言中没有try...catch来捕获错误，一旦触发panic就会导致程序崩溃。在Go中是通过recover让程序恢复。

⚠️ 值的注意的是recover()必须在延迟函数(defer)中有效。

*/

func main() {
	// 定义一个匿名延迟函数
	defer func() {
		err := recover()
		msg := fmt.Sprintf("err信息: %v", err)
		if err != nil {
			// 程序触发panic时，会被这里捕获
			fmt.Println("err = ", err)
			fmt.Println("msg = ", msg)
		}

		fmt.Println("err = ")
	}()
	// 定义一个数组
	testPanic("请求失败")
}

// 故意抛出panic
func testPanic(err string) {
	panic("错误信息" + err)
}
