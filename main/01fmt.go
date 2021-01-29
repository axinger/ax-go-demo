package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	//占位符
	/*
		Print:   输出到控制台,不接受任何格式化操作
		Println: 输出到控制台并换行
		Printf : 只可以打印出格式化的字符串。只可以直接输出字符串类型的变量（不可以输出别的类型）
		Sprintf：格式化并返回一个字符串而不带任何输出
		Fprintf：来格式化并输出到 io.Writers 而不是 os.Stdout

	*/

	fmt.Println("字节长度 = ",len("汉字"))
	fmt.Println("子串长度 要减 1= ",strings.Count("12","")-1)

	fmt.Println("子串长度3 = ",utf8.RuneCountInString("汉字"))

	var num = 3
	//var num = true
	// %T 查看类型
	fmt.Printf("查看类型 %T \n",num)
	// %v 查看值
	fmt.Printf("查看值 %v \n",num)
	map1 := make(map[string]interface{})
	map1["name"] = "jim"
	fmt.Printf("查看值#v %#v \n",map1)
	fmt.Printf("查看值+v %+v \n",map1)
	fmt.Printf("查看值v %v \n",map1)

	fmt.Printf("查看值 %+v \n",num)

	fmt.Println("Println查看值",num)


	// 非整数错误
	fmt.Printf("二十进制 %b \n",num)
	// 非整数错误
	fmt.Printf("十进制 %d \n",num)
	fmt.Printf("八进制 %o \n",num)
	fmt.Printf("十六进制 %x \n",num)

	var str = "aaa"
	fmt.Printf("字符串 %s \n",str)

	fmt.Printf("查看值 %v \n",12)


	/// Println
	fmt.Println("Println逗号拼接","aa",12)

	fmt.Sprintf("Sprintf")
}
