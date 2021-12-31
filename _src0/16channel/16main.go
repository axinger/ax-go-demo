package main

import (
	"fmt"
	"sync"
	"time"
)

var c1 chan int

var wg sync.WaitGroup
var once sync.Once

func main() {
	// 传递int类型的通道
	// 指针类型,需要初始化

	//test1()
	//test2()
	//c3 := make(chan int, 3)
	//c3<-1
	//c3<-2
	//test3(c3)
	//
	//test11()

	workerDo()
}

/// 这个为什么可以,异步为什么能添加数据
func test11() {
	ch := make(chan int) //创建一个无缓存channel

	//新建一个goroutine
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i //往通道写数据
		}
		//不需要再写数据时，关闭channel
		close(ch)
		//ch <- 666 //关闭channel后无法再发送数据

	}() //别忘了()

	for num := range ch {
		fmt.Println("num = ", num)
	}

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
	//// 要先关闭,才能range 遍历
	//close(c2)
	//
	////}()
	//for data := range c2 {
	//	fmt.Println("c2 = ", data)
	//}

	max := len(c2)
	for i := 0; i < max; i++ {
		fmt.Println("c2 = ", <-c2)
	}

}

// 单向通道,用于参  两种形式
// chan<- ,添加值,不能取值
// <-chan 只能取值,不能添加值
func test3(c <-chan int) {

	//c <- 2
	//c <- 3
	//fmt.Println("len(c) = ",len(c))
	// len 取值一次,就会减少一个,所以要取值
	max := len(c)
	for i := 0; i < max; i++ {
		fmt.Println("test3 = ", <-c)
	}

	//for i := range c {
	//	fmt.Println("test3 = ",i)
	//}
	//close(c)
	//x:= <-c
	//fmt.Println(x)
	//fmt.Println("test3 = ",c)
}

func worker(id int, jobs <-chan int, res chan<- int) {
	for job := range jobs {
		fmt.Println("id = ", id, "开始 = ", job)
		time.Sleep(time.Second)
		fmt.Println("id = ", id, "结束 = ", job)
		res <- job * 2
	}
}

func workerDo() {
	jobs := make(chan int, 100)
	res := make(chan int, 100)
	// 开启3个 goroutine
	for i := 0; i < 3; i++ {
		go worker(i, jobs, res)
	}

	// 发送添加值
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)

	//取值
	for i := 0; i < 5; i++ {
		<-res
		//fmt.Println("res = ",<-res)
	}

}
