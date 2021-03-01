package main

import (
	"container/list"
	"fmt"
	"sort"
)

func main() {

	// 数组长度固定,一般常用切片
	/// 初始值长度 ...
	var list1 = [...]string{"1", "2"}

	//list1 = [2]string{"1","2"}

	list1[1] = "33"
	fmt.Println(list1)
	fmt.Println(len(list1))
	list2 := [5]int{0: 1, 4: 5}
	fmt.Println(list2)
	for i := range list2 {
		fmt.Println(i)
	}

	for i, i2 := range list2 {
		fmt.Println(i, i2)
	}

	// 切片 slice
	list3 := []int{1, 2}
	//append(list3, "3")
	/// 直接追加 多个
	list3 = append(list3, 3, 4)

	/// 追加切片 ... 表示拆开
	list3 = append(list3, []int{5, 6}...)

	if list3 != nil {
		fmt.Println("list3 != nil")
	}
	fmt.Println("切片赋值list3 = ", list3)
	fmt.Println("cap策略不同", cap(list3))
	fmt.Println("len", len(list3))

	/// 数组转化切片
	arr1 := [...]int{1, 2, 3, 4, 5, 6}
	list4 := arr1[1:3]
	list4 = arr1[:3]
	list4 = arr1[1:]
	list4 = arr1[:]
	list4 = append(list4, 7)
	fmt.Println("数组转化切片list4", list4)
	/// makr 开辟空间
	var list5 = make([]int, len(list4))
	//var list5 = []int
	//var	list5 = make([]int, len(list4), cap(list4))
	//	var	list5 = make([]int, cap(list4), len(list4))
	copy(list5, list4)
	fmt.Println("切片copy", list5)

	list5 = append(list5[:1], list5[5:]...)
	fmt.Println("切片没有专门删除方法 ,切片删除", list5)

	// 切片不存值,底层数组存值
	/// list6 底层数组为 arr6
	arr6 := [...]int{1, 3, 5, 7, 9}
	list6 := arr6[:]
	fmt.Println("arr6 = ", arr6)

	fmt.Println("list6 = ", list6)

	//list6 = append(list6, 4)
	//fmt.Println("append rr6 = ",arr6)
	//fmt.Println("append list6 = ",list6)

	list6 = append(list6[:1], list6[3:]...)
	fmt.Println("删除 arr6 = ", arr6)   // 1,5,5 为什么? 把第二个切了 1,5 cap扩大,追加最后几位补齐
	fmt.Println("删除 list6 = ", list6) // 1,5

	//var list7 = make([]int,5)
	//	var list7 = make([]int,5,3)
	//	var list7 = make([]int,5,5)
	var list7 = make([]int, 5, 6)
	fmt.Println("list7", list7)
	for i := 0; i < 10; i++ {
		list7 = append(list7, i)
	}
	fmt.Println("list7", list7) // 5个0 ,012345...

	var list8 = make([]int, 9)

	fmt.Println("list8", list8)
	for i := 0; i < len(list8); i++ {
		list8[i] = len(list8) - i - 1
	}
	fmt.Println("for 初始化切片list8", list8) // 5个0 ,012345...

	sort.Ints(list8)
	fmt.Println("排序 list8", list8)
	/// make 只能用于切片 map chan 内存创建
	/// new 很少用,基本数据类型申请内存
	var aint *string = new(string)
	fmt.Println(aint)


	list21 := list.New()
	fmt.Println("list21 = ",list21)
	list21.PushBack("A")
	list21.PushBack("B")
	fmt.Println("list21 = ",list21)


	fmt.Printf("%v \n", list21)

	fmt.Println("Front = ",list21.Front().Next())
	for item := list21.Front(); item != nil; item = item.Next() {
		fmt.Println("item = ",item.Value)
		//fmt.Println(item.Value.([]byte))
	}
}
