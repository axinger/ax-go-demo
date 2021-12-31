package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var x = 0

// 互斥锁
var lock sync.Mutex

func add() {
	defer wg.Done()
	for i := 0; i < 5000; i++ {
		x = x + 1
	}

}

// 竞争资源,x != 1000
func test1() {

	wg.Add(2)
	go add()
	go add()

	wg.Wait()
	fmt.Println("执行完成===============", x)

}

// 使用互斥锁,
func add2() {
	defer wg.Done()
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
}

// 互斥锁
func test2() {

	wg.Add(2)
	go add2()
	go add2()

	wg.Wait()
	fmt.Println("执行完成===============", x)
}
func main() {
	//test1()
	test2()
}
