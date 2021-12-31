package main

import (
	"flag"
	"fmt"
)

// 获取命令行参数
func main() {

	// 不太好用,空格会被解析
	//fmt.Println("os.Args::", os.Args)
	//
	//// 用于脚本获取参数
	//fmt.Println("flag.Arg::", flag.Args())
	//fmt.Println("flag.Arg::", flag.Arg(1))

	// 创建标志位  -help   获得提示
	// 返回指针
	name := flag.String("name", "jim", "请输入姓名")

	// 返回指针
	age := *flag.String("age", "10", "年龄")

	//指针传值
	var height string
	flag.StringVar(&height, "height", "170cm", "身高")

	//一定要先解析变量
	flag.Parse()

	fmt.Printf("name: 类型:%T,取值:%v\n", name, *name)
	fmt.Printf("age: 类型:%T,取值:%v\n", age, age)
	fmt.Printf("height: 类型:%T,取值:%v\n", height, height)

	fmt.Println("没有占位的内容flag.Args:", flag.Args())
	fmt.Println("没有占位的个数flag.NArg:", flag.NArg())

}
