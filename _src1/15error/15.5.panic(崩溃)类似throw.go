package main

import "fmt"

/*

panic让当前的程序进入恐慌，中断程序的执行。panic()是一个内建函数，可以中断原有的控制流程。其功能类似于中的throw。

*/

func throw(msg string, code int) {
	if code == 0 {
		panic("报错信息: " + msg)
	}
	fmt.Println("正确输出：" + msg)
}
func main() {
	// 测试访问数组下标越界

	// 测试-> 访问未初始化的指针或 nil 指针
	// 1.定义一个指针类型（默认值是nil）
	//var b *int
	//fmt.Println(b)
	//// 2.访问nil的指针
	//fmt.Println(*b)
	throw("请求成功!", 1)
	throw("请求失败!", 0)

}
