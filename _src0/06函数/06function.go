package main

import "fmt"

//命名返回值
func test1(x, y int, a, b string) (res int) {
	res = 3
	return

}

// 多个返回值
func test1_1(x, y int, a, b string) (res1 int,res2 int) {

	return 1,2

}


// 可变参数
func test2(a int, b ...int) {
	// b是切片类型
	fmt.Println("a = ", a, "b = ", b)

}
func test3(b ...int) {
	// b是切片类型
	fmt.Println("b = ", b)

}

// go 的return 不是原子操作,在底层分为两步执行
// 第一步 返回值赋值
// 第二步 真正的RET返回
// 函数中如果存在 defer ,那么defer执行的时机是在第一步和第二步之间
func test4(a int) int {
	res := a
	defer func() {
		res++
		fmt.Println("res = ", res)
	}()

	return res
}

//
func test5(a int) (res int) {
	defer func() {
		res++
		// 这里 res++
		//fmt.Println("test5 = ", res)
	}()

	/// 这里会改变 返回值 res =a
	/// 最后参会 res++
	return a
}

func test6(a int) (res int) {
	//x :=5

	defer func() {
		res++
		// 这里 res++
		//fmt.Println("test5 = ", res)
	}()

	/// 这里会改变 返回值 res =a
	/// 最后参会 res++
	return a
}

func test7(x func() int) {
	res := x()
	fmt.Println("test7 函数参数返回值= ", res)
}

func test8(x func() int) func(int, int) int {
	res := x()
	fmt.Println("test8 函数参数返回值= ", res)

	return func(i int, i2 int) int {
		return i + i2
	}

}

func test9(x func(a int) int) {
	res := x(9)
	fmt.Println("test9 函数参数返回值= ", res)
}

var test10 = func(a, b int) {
	fmt.Println("test10 = ", a, b)

	// 函数内部使用匿名函数
	res := func(a, b int) {
		fmt.Println("test10 函数内部使用匿名函数= ", a, b)

	}
	res(10,11)
}

func main() {
	// 推迟
	//defer test3(-1)
	test2(1, 2)
	test3(3, 4, 5)
	fmt.Println("test4 = ", test4(10))
	fmt.Println("test5 = ", test5(10))
	fmt.Println("test6 = ", test6(10))

	// 无法预期效果 defer 执行的变量,不是后面赋值结果
	///defer 用于函数结束之前 释放资源
	//var age = "aa"
	//defer fmt.Println("age = ",age)
	//age = "jim"
	//fmt.Println("age2 = ",age)

	test7(func() int {
		fmt.Println("test7 函数参数")
		return 5
	})

	f := test8(func() int {
		return 8
	})
	res8 := f(8, 9)
	fmt.Println("res8 = ", res8)

	test9(func(a int) int {
		return a + 10
	})
	test10(10, 11)


// 闭包
}
