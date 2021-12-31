package main

import (
	"fmt"
	"strconv"
	"time"
)

/*

使用go关键字创建Goroutine时，被调用的函数往往没有返回值，
⚠️ 如果有返回值也会被忽略。如果需要在Goroutine中返回数据，必须使用channel，通过channel把数据从Goroutine中作为返回值传出。
*/

func main() {
	go echoNum("张三")
	go echoNum("李四")
	go echoNum("王五")

	// 多个Goroutine随机调度，打印的结果是数字与字母交叉输出
	go func(who string) {
		for i := 1; i < 3; i++ {
			fmt.Println(who + " " + strconv.Itoa(i))
		}
	}("张三")
	go func(who string) {
		for i := 1; i < 3; i++ {
			fmt.Println(who + " " + strconv.Itoa(i))
		}
	}("李四")
	go func(who string) {
		for i := 1; i < 3; i++ {
			fmt.Println(who + " " + strconv.Itoa(i))
		}
	}("王五")
	// 睡眠1秒
	time.Sleep(1 * time.Second)
	fmt.Println("运行结束")

	// 睡眠1秒
	time.Sleep(1 * time.Second)
	fmt.Println("运行结束")
}

// 报数
func echoNum(name string) {
	for i := 1; i < 3; i++ {
		//fmt.Println(name + " "+ strconv.Itoa(i))
		fmt.Println(name, ">>>>", i)
	}
}

/** 输出
  张三 1
  张三 2
  王五 1
  李四 1
  李四 2
  王五 2
  运行结束
*/
