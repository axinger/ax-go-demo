package main

import "fmt"

// 动物接口
type Animal interface {
	eat(name string) string

	run(name string)
}

type Dog struct {
	name string
}

func (d Dog) eat(s string) string {

	return "aaa"
}

type Cat struct {
	name string
}

func (c Cat) eat(s string) string {
	fmt.Println("cat,实现接口")
	return "aaa"
}

func  (c Cat)run(string2 string)  {

}

//func (a Animal)jump()  {
//
//}

// 接口作为参数
func jump2(a Animal) {

}

func main() {

	var a1 Animal
	c1 := Cat{
		name: "tom",
	}
	a1 = c1

	a1.eat("cat")

}
