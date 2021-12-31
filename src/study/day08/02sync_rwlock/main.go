package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg   sync.WaitGroup
	x    = 0
	lock sync.Mutex
	// 读写互斥锁
	rwLock sync.RWMutex
)

func read() {
	defer wg.Done()
	//lock.Lock()
	rwLock.RLock() //加读锁
	fmt.Println("读操作...", x)
	time.Sleep(time.Millisecond)
	//lock.Unlock()
	rwLock.RUnlock()
}

func write() {
	defer wg.Done()
	//lock.Lock()
	rwLock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 5)
	//lock.Unlock()
	rwLock.Unlock()
}

// 使用互斥锁
func test1() {

	start := time.Now()
	for i := 0; i < 100; i++ {
		go write()
		wg.Add(1)
	}

	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}

	wg.Wait()
	enb := time.Now()
	fmt.Println("耗时:", enb.Sub(start))

}

// 读写互斥,读的次数大于写的次数
func main() {
	test1()
}
