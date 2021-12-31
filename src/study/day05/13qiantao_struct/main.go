package main

import "fmt"

type info struct {
	province string
	city     string
}

type person struct {
	name string
	age  int
	info // 匿名字段 嵌套 没有什么特别用处
}

func main() {

	p := person{
		name: "jim",
		age:  10,
		info: info{province: "江苏省", city: "南京市"},
	}
	fmt.Println("p = ", p, p.province, p.info.province)

	p2 := person{}
	fmt.Println("没有初始化,的值 p2.info = ", p2.info)
	fmt.Println("嵌套结构体,没有初始化,p2 = ", p2, p2.info, p2.province, p2.info.province)

}
