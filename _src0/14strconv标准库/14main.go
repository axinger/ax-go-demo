package main

import (
	"fmt"
	"strconv"
)

// 类型转换
func main() {

	i := int32(97)
	fmt.Println("i = ",i)
	str1 :=string(i)
	fmt.Println("str = ",str1)
	str3,_:= strconv.ParseInt("97", 10, 64)
	fmt.Println("str3 = ",str3)

	// base 当前子串的 进制 16进制 转 10进制
	str4,_:= strconv.ParseInt("a", 16, 64)
	fmt.Println("str4 = ",str4)

	// a to i c语言中的 array to int
	str5,_ := strconv.Atoi("11")
	fmt.Println("str5 = ",str5)
	str6 := strconv.Itoa(010)
	fmt.Println("str6 = ",str6)

}
