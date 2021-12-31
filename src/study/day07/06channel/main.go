package main

import (
	"fmt"
	"sync"
)

// 定义通道 关键字 类型
var c1 chan int

func main() {

	//test1()

	test2()
}

var waitGroup sync.WaitGroup

func test1() {

	waitGroup.Add(1)

	// 一定要初始化,无缓冲区
	c1 = make(chan int)

	// 开启线程,准备接收值
	go func() {
		defer waitGroup.Done()
		// 接收值
		num := <-c1
		fmt.Println("通道接收值:", num)
	}()

	// 模拟 另一个线程 发送值
	go func() {
		// 发送值
		c1 <- 1
	}()

	//close(c1)

	waitGroup.Wait()
	fmt.Println("执行完成===================")
}

func test2() {
	fmt.Println("有缓冲区,未满,不阻塞===================")
	// 一定要初始化,有缓冲区
	c1 = make(chan int, 1)

	// 发送值
	c1 <- 1
	//c1 <- 2 // 满了,就会阻塞

	// 接收值
	//num := <-c1
	//fmt.Println("通道接收值:", num)

	close(c1)

	fmt.Println("执行完成===================")
}
