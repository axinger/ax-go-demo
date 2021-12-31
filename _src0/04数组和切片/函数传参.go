package main

import "fmt"
func testSlice(slice []int)  {
	// 修改切片的值
	slice[0] = 100
}
func main() {
	// 创建一个切片
	slice := []int{1,2,3,4}
	fmt.Printf("变量slice  类型: %T 内存地址: %p 值:%v \n",slice,&slice,slice)
	// 调用函数修改
	testSlice(slice)
	fmt.Printf("调用函数后,变量slice  类型: %T 内存地址: %p 值:%v \n",slice,&slice,slice)
}
/**输出
变量slice  类型: []int 内存地址: 0xc00000c080 值:[1 2 3 4]
调用函数后,变量slice  类型: []int 内存地址: 0xc00000c080 值:[100 2 3 4]
*/
