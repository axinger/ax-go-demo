package main

import (
	"fmt"
	"runtime"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	//runtime.GOMAXPROCS(1) // 1个线程,就会串行
	//runtime.GOMAXPROCS(4)
	fmt.Println("cpu数据:", runtime.NumCPU())
	waitGroup.Add(2)
	go a()
	go b()
	waitGroup.Wait()
	fmt.Println("执行结束===================")
}

func a() {
	defer waitGroup.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("a===", i)
	}
}
func b() {
	defer waitGroup.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("b===", i)
	}
}
