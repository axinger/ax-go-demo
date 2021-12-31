package main

import "fmt"

func main() {
	// 测试
	defer test1()

	// 匿名函数
	defer func() {
		fmt.Println("匿名函数defer3....")
	}()
	// ⚠️ 多个这些defer语句会按照逆序执行

	// 函数正常处理代码
	for i := 1; i < 4; i++ {
		fmt.Println(i)
	}
}

// 无参数的函数
func test1() {
	fmt.Println("函数defer1....")
}

/**输出
  1
  2
  3
  匿名函数defer3....
  函数defer2....a=40 b= 60 a+b=100
  函数defer2....a=4 b= 6 a+b=10
  函数defer1....
*/
