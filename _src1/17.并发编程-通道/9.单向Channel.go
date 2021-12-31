package main

import (
	"fmt"
	"time"
)

func main() {
	workerDo()

	//{
	//	// 创建一个整数型chan
	//	intChan := make(chan int)
	//	go writeChan(intChan)
	//	go readChan(intChan)
	//	time.Sleep(50 * time.Millisecond)
	//	fmt.Println("运行结束")
	//
	//}

}

/// ch <- chan T
func worker(id int, jobs <-chan int, res chan<- int) {
	for job := range jobs {

		fmt.Println("jobs接收值", job, "id = ", id)

		//fmt.Println("id = ", id, "开始 job= ", job)
		time.Sleep(5 * time.Second)
		//fmt.Println("id = ", id, "结束 job= ", job)
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
	for i := 10; i < 15; i++ {
		fmt.Println("jobs发送值", i)
		jobs <- i
	}
	close(jobs)

	//取值
	for i := 0; i < 5; i++ {
		val := <-res
		fmt.Println("res = ", val)
	}

	//for re := range res {
	//	fmt.Println("res取值 = ", re)
	//}
}

// 定义只读通道的函数
func readChan(ch <-chan int) {
	for data := range ch {
		fmt.Printf("读出数据: %v \n", data)
	}
}

// 定义只写通道的函数
func writeChan(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i
		fmt.Printf("写入数据: %v \n", i)
	}
	close(ch)
}
