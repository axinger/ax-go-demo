package main

import "fmt"

func main() {



	//使用make()函数来创建切片
	var slice1 = make([]int,3)
	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 4
	fmt.Printf("通过make关键字[slic1] 类型:%T 值:%v \n",slice1,slice1)


	// 直接初始化使用
	//slice2 := []int{}
	var slice2 []int
	for i := 0; i < 5; i++ {
		slice2 = append(slice2,i)
	}
	fmt.Printf("直接初始化使用[slic2] 类型:%T 值:%v \n",slice2,slice2)


	// 定义一个数组
	arr := [...]int{1,2,4,6,8,10,12,14,16}
	// 从索引0开始截取到索引为4(不包括4)
	slice3 := arr[0:4]
	fmt.Printf("从索引0开始截取到索引为4(不包括4)[slice3] 类型:%T 值:%v \n",slice3,slice3)
	// 索引从0开始时，0也可以省略不写
	slice33 := arr[:4]
	fmt.Printf("从索引0开始截取到索引为4(不包括4)[slice33] 类型:%T 值:%v \n",slice33,slice33)
	// 从索引4开始截取到最后
	slice4 := arr[5:]
	fmt.Printf("从索引5开始截取到最后[slice4] 类型:%T 值:%v \n",slice4,slice4)


	/// 3.切片删除
	/// 3.1 删除第一个元素
	// 定义一个切片
	slice := []int{1,2,3,4,5,6}
	fmt.Printf("切片slice--> 值:%v len: %d cap: %d\n",slice,len(slice),cap(slice))
	// 删除第一个元素
	slice = slice[1:]
	fmt.Printf("删除一个元素后--> 值:%v len: %d cap: %d\n",slice,len(slice),cap(slice))


	/// 3.2 删除最后一个元素
	fmt.Printf("切片slice--> 值:%v len: %d cap: %d\n",slice,len(slice),cap(slice))
	// 删除最后一个元素
	slice = slice[:len(slice)-1]
	fmt.Printf("删除一个元素后--> 值:%v len: %d cap: %d\n",slice,len(slice),cap(slice))


	// 3.3 删除指定位置元素
	// 切片没有专门的删除方法,
	index := 2
	// 指定索引前面一部分
	front := slice[:index]
	// 指定索引后面一部分
	behind := slice[index+1:]
	// 拼起来
	slice =  append(front,behind...)
	fmt.Printf("删除索引%d后--> 值:%v len: %d cap: %d\n",index,slice,len(slice),cap(slice))

}
