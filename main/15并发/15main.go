package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

// goroutine 和 channel 协调
// main 就是一个goroutine
func main() {

	// 函数 前面 加 go 就创建 了  goroutine 并发
	// 并发 一个人 一段时间同时做多个事情
	//并行  多个人做多个事情
	go hello()
	fmt.Println("main=====")

	for i := 0; i < 10; i++ {
		//匿名
		// 这里因为作用域.可能多个
		//go func() {
		//	fmt.Println("hello2=====",i)
		//}()

		// 用参数,避免上面的问题
		//go func(i int) {
		//	rand.Seed(time.Now().UnixNano())
		//	time.Sleep(time.Millisecond*(time.Duration(rand.Intn(500))))
		//	fmt.Println("hello2=====", i)
		//}(i)
		// 添加一个,和donw 减一个 匹配
		wg.Add(1)
		go group(i)
	}

	//time.Sleep(time.Second)
	wg.Wait()
}

func hello() {
	fmt.Println("hello===")
}
func group(i int) {
	// 减少一个 匹配
	//defer wg.Done()
	// Done 就是 -1
	defer wg.Add(-1)

	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Millisecond * (time.Duration(rand.Intn(500))))
	fmt.Println("hello2=====", i)

}

// 随机数
func rand2() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		intn := rand.Intn(100)
		fmt.Println("随机数 = ", intn)
	}
}
