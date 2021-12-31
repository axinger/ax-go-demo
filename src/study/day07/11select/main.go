package main

import "fmt"

func main() {

	test1()
}

// 多路复用
func test1() {
	// select 随机取值,可以同时取值,发送值

	//ch := make(chan int, 1) // 1 必然会交替执行
	ch := make(chan int, 3)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println("随机取值........", x)

		case ch <- i:
			fmt.Println("随机发送值=====")
		default:
			fmt.Println("都无法执行>>>>>>>>>>>>")
		}
	}
}
