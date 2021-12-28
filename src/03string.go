package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符串操作
	// 1.拼接
	name := "jim"
	name2 := "tom"

	name = "jim2"
	fmt.Println(name + name2)
	// for range 键值遍历
	for index, i := range name {
		//fmt.Println("子串for", i,"index = ",index)
		fmt.Printf("子串for %d %c \n", index,i)
	}

	// mt.Sprintf 打印含有返回值
	name3 := fmt.Sprintf("%s%s", name, name2)
	fmt.Println(name3)
	/// 分割 成数组
	name4 := strings.Split("1,2,3,", ",")
	fmt.Println("name4 = ", len(name4))

	fmt.Println("Contains = ", strings.Contains("jim", "ji"))
	/// 切片可变的,数组不可变的
	list1 := []string{"1", "2"}
	list1 = append(list1, "3")

	fmt.Println("拼接 = ", strings.Join(list1, ","))

	/// rune 类型的切片,非英文都是rune类型 实际就是 []int32 类型
	name5 := []rune(name)
	name5[0] = 't'
	fmt.Println("修改后 = ", string(name5))

	fmt.Printf("类型 name = %T name5 = %T\n", name, name5)
	{
		fmt.Printf("强转 = %T \n", 10)
		fmt.Printf("强转 = %T \n", float64(10))
	}

}
