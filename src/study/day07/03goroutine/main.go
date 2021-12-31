package main

import (
	"fmt"
)

func main() {

	test0()
}

func test0() {

	/*
		go+函数:开启一个单独的goroutine执行函数
		 多个 goroutine ,逆序执行
	*/
	go test01()
	//go func() {
	//	fmt.Println("匿名函数...........")
	//}()
	for i := 0; i < 100; i++ {
		// 闭包会强引用 i
		go func(j int) {
			fmt.Println("匿名函数...........", j)
		}(i)

	}

	fmt.Println("main...........")

	//time.Sleep(5 * time.Second)
}

func test01() {

	fmt.Println("test1...........")
}
