package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// int 并发时候,需要加锁, 直接用原子操作就行了

var num int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {

	//lock.Lock()
	//num = num + 1
	//lock.Unlock()

	// 使用原子性
	atomic.AddInt64(&num, 1)

	/*

	   num1  :=  int64(43)
	   	atomic.AddInt64(&num1,1)*/

	wg.Done()

}

func test1() {
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go add()
	}
	wg.Wait()
	fmt.Println("并发执行完成:", num)
}
func main() {
	//test1()

	// 比较并交换,相同时,用新值替换原值
	num1 := int64(1)
	swapInt64 := atomic.CompareAndSwapInt64(&num1, 1, 200)
	fmt.Println("num1:", num1, "swapInt64:", swapInt64)
}
