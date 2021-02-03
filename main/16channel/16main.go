package main

import (
	"fmt"
	"sync"
)

var c1 chan int

var wg sync.WaitGroup
var once sync.Once

func main() {
	// 传递int类型的通道
	// 指针类型,需要初始化

	//test1()
	test2()

}
func test1() {
	c1 = make(chan int) // 无缓存,会卡,用线程控制
	//var c1 chan int = make(chan int,16) //缓冲区通道,不会卡直接运行
	fmt.Println(c1)

	// 定义取值
	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := range c1 {
			fmt.Println("循环取值 = ", i)
		}

		//x := <-c1
		//fmt.Println("c1 = ", x)
	}()
	// 发送值
	c1 <- 2
	c1 <- 3
	c1 <- 4
	close(c1)

	wg.Wait()
	//接收
	//x := <-c1

	//fmt.Println(x)

}

func test2() {
	//defer once.Do(func() {
	//	close(c2)
	//})
	//缓冲区通道,不会卡直接运行
	c2 := make(chan int, 6)
	//defer close(c2)
	//go c2Add()
	//go func() {
	c2 <- 1
	c2 <- 2
	// 要先关闭,才能range 遍历
	close(c2)

	//}()
	for data := range c2 {
		fmt.Println("c2 = ", data)
	}

}
