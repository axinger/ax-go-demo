package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	testWait()
}

var waitGroup sync.WaitGroup

// 随机数
func rand2() {
	// 随机种子,防止 随机数 每次都一样
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		num := rand.Intn(100)
		fmt.Println("随机数 = ", num)
	}
}

func testWait() {

	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go f1(i)
	}
	waitGroup.Wait()
	fmt.Println("执行完成===================")
}

func f1(i int) {
	defer waitGroup.Done()

	// 随机睡眠
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	fmt.Println("随机睡眠:", i)

}
