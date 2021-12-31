package main

import (
	"fmt"
	"strconv"
)

func main() {

	i := 97
	s := string(i)
	fmt.Printf("int 强转 string 类型:%T ,值: %v \n", s, s)
	sprintf := fmt.Sprintf("%d", i)
	fmt.Printf("fmt.Sprintf 类型:%T ,值: %v \n", sprintf, sprintf)
	// int 强转 string 类型:string ,值: a ,string 当时ASC码,所以不要用string(i) 进行强转

	// 10 表示十进制, 32,64 int32 int64
	parseInt, error := strconv.ParseInt("97", 10, 0)
	if error != nil {

		return
	}
	fmt.Printf("字符串解析对应数据 类型:%T ,值: %#v \n", parseInt, parseInt)

	// 字符串 转 int
	num2, error := strconv.Atoi("97")
	if error != nil {
		return
	}
	fmt.Printf("字符串 转 int Atoi 类型:%T ,值: %#v \n", num2, num2)

	intVal := int(97)
	str2 := strconv.Itoa(intVal)
	fmt.Printf("int 转 字符串 Itoa 类型:%T ,值: %#v \n", str2, str2)

	boolStr := "true"
	parseBool, error := strconv.ParseBool(boolStr)
	if error != nil {
		return
	}
	fmt.Printf("字符串 转 bool ParseBool 类型:%T ,值: %#v \n", parseBool, parseBool)

}
