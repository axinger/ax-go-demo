package main

import "fmt"

// 动物接口
type Animal interface {
	// 接口嵌套
	eater
	jumper
	//eat(name string) string
	//run(name string)
}

// 空接口,无方法,任何都可以接收,不实现方法 interface{} map中使用

func show(a interface{}) {
	//fmt.Printf("空接口 %T", a)
	//st1, ok := a.(string)
	//if ok {
	//	fmt.Println("类型转换 拍 = ",st1,"拍 = ",ok)
	//}

	switch t := a.(type) {
	case string:
		fmt.Println("拍 string = ", t)
	case int:
		fmt.Println("拍 int = ", t)
	}

}

// 这个不行
//func (a1 interface{})show1(a interface{})  {
//	fmt.Printf("空接口 %T",a)
//}

// 接口后面+er后缀
type eater interface {
	eat(name string)
}

type jumper interface {
	jump(name string)
}

type mover interface {
	move(name string)
}

type Cat struct {
	name string
}

// 实现接口,值接收者
func (c Cat) eat(s string) {
	fmt.Println("cat,实现接口")
}

func (c Cat) jump(string2 string) {

}

// 多实现接口
func (c Cat) move(string2 string) {
	fmt.Println("cat,多实现接口 = move")
}

/**

// 实现接口,指针接收者
func (c *Cat) eat(s string) string {
	fmt.Println("cat,实现接口")
	return "aaa"
}

func  (c *Cat)run(string2 string)  {

}

*/

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
	c1.move("c1")
	show("jim")
	show(2)
	fmt.Printf("接口得到的类型,得到赋值类型 = %T \n", a1)

	//c2 :=& Cat{
	//	name: "tom_c2",
	//}
	//a1 = c2
	//fmt.Println("指针接口实现接收者 = ",a1)
}
