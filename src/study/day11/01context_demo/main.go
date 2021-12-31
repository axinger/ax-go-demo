package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test1(ctx context.Context) {
	defer wg.Done()

	go test2(ctx)
ForLoop:
	for {
		fmt.Println("test1............")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			// 指定 跳转到范围
			break ForLoop
		default:

		}
	}
}

func test2(ctx context.Context) {
	defer wg.Done()

ForLoop:
	for {
		fmt.Println("test2............")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			// 指定 跳转到范围
			break ForLoop
		default:

		}
	}
}

func main() {
	// context.Background() 第一层的
	ctx, cancelFunc := context.WithCancel(context.Background())

	wg.Add(1)
	go test1(ctx)

	time.Sleep(time.Second * 5)
	// 通知通道取消
	cancelFunc()
	wg.Wait()
	// 通知子goroutine退出
}
