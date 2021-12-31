package main

import (
	"fmt"
	"sync"
)

func main() {

	a := make(chan int, 100)
	b := make(chan int, 100)
	waitGroup.Add(3)

	go test1(a)
	go test2(a, b)
	go test2(a, b)
	test3(b)
	waitGroup.Wait()
	fmt.Println("执行完成===================")
}

//var ch1 = make(chan int, 6)
var waitGroup sync.WaitGroup
var once = sync.Once{}

// ch1 传值
func test1(ch1 chan int) {
	defer waitGroup.Done()

	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

// ch1 读取数据,然后传值 ch2
func test2(ch1 chan int, ch2 chan int) {
	defer waitGroup.Done()
	//for i := range ch1 {
	//	ch2 <- i * 10
	//}

	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * 10
	}

	// 执行一次,必须这样写,不能链式,必须全局
	//o := sync.Once{}
	once.Do(
		func() {
			close(ch2)
		})

}

// ch2 读取数据
func test3(ch2 chan int) {
	for i := range ch2 {
		num := i
		fmt.Println("ch2取值:", num)
	}
}
