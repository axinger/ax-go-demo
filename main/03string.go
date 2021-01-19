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
	fmt.Println(name + name2)

	// mt.Sprintf 打印含有返回值
	name3 := fmt.Sprintf("%s%s", name, name2)
	fmt.Println(name3)
	/// 分割 成数组
	name4 := strings.Split("1,2,3,", ",")
	fmt.Println("name4 = ", len(name4))

	fmt.Println("Contains = ", strings.Contains("jim", "ji"))
	/// 切片可变的,数组不可变的
	list1 := []string{"1","2"}
	list1 = append(list1, "3")

	fmt.Println("拼接 = ",strings.Join(list1,","))

}