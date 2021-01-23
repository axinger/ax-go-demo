package main

import "fmt"

func test11(f func()) {
	f()
}

func test12(x, y int) {
	fmt.Println(x + y)
}

func test13(x int) func(int) int {

	// 使用了 x 这个变量,就是闭包
	return func(y int) int {

		return x + y
	}

}

// 闭包
func main() {

	test11(func() {
		test12(1, 2)
	})

	res13 := test13(2)(10)
	fmt.Println("res13 = ", res13)

}
