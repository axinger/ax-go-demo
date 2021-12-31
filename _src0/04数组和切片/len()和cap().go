package main

import "fmt"
/**
切片的长度是切片中元素的数量。
切片的容量是从创建切片的索引开始的底层数组中元素的数量。
切片可以通过len()方法获取长度，可以通过cap()方法获取容量。数组计算cap()结果与len()相同。
*/
func main() {

	// 创建一个切片
	slice := []int{1,2,4,6,8,10,12,14,16}
	fmt.Printf("切片slice， len:%d cap:%d \n",len(slice),cap(slice))

	// 清空切片slice赋值给slice1
	slice1 := slice[0:0]
	fmt.Printf("切片slice1， len:%d cap:%d \n",len(slice1),cap(slice1))

	// 使用make指定容量创建切片
	var slice2 = make([]int,4,8)
	slice2[1] = 1
	slice2[3] = 3
	fmt.Printf("切片slice2， len:%d cap:%d \n",len(slice2),cap(slice2))

	// 数组计算cap()结果与len()相同
	arr := [...]int{1,2,4,6,8,10,12,14,16}
	fmt.Printf("数组arr， len:%d cap:%d \n",len(arr),cap(arr))

}
