package main

import "fmt"

func test11(f func()) {
	f()
}

func test12(x, y int) {
	fmt.Println(x + y)
}

func test13(x int) func(y int) int {

	// 使用了 x 这个变量,就是闭包
	return func(y int) int {

		return x + y
	}

}

func counterByClosure() func(num int) {
	sum := 0
	return func(num int) {
		fmt.Printf("num: %d 相加前 sum = %d  相加后 sum =  %d \n",num,sum,sum+num)
		sum += num
	}
}

// 闭包
func main() {

	test11(func() {
		test12(1, 2)
	})

	res13 := test13(2)(10)
	fmt.Println("res13 = ", res13)

	counterByClosure()(7)


	// 初始化闭包函数
	counterFunc := counterByClosure()
	fmt.Println(" ---------  再从1加到10 -------- ")
	// 从1加到10
	for i := 1; i<= 10; i++ {
		counterFunc(i)
	}
	fmt.Println(" ---------  再从10加到20 -------- ")
	// 再从10加到20
	for i := 10; i<= 20; i++ {
		counterFunc(i)
	}

}
