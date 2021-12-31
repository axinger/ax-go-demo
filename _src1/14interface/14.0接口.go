package main

import "fmt"

/*
虽然Go语言没有继承和多态，但是Go语言可以通过匿名字段实现继承，通过接口实现多态。

在Go语言中，接口是一组方法签名。接口指定了类型应该具有的方法，类型决定了如何实现这些方法。
当某个类型为接口中的所有方法提供了具体的实现细节时，这个类型就被称为实现了该接口。
接口定义了一组方法，如果某个对象实现了该接口的所有方法，则此对象就实现了该接口。


type 接口名称 interface {
    Method1([参数列表]) [返回值列表]
    Method2([参数列表]) [返回值列表]
    ...
}

// 定义一个接口
type Bird interface {
    fly() // 无参数无返回值方法
    eat(string2 string) // 有参数无返回值方法
    walk(string2 string) string // 有参数有返回值方法
}


*/

// 定义一个鸟类接口
type Birder interface {
	fly()
	eat(food string)

	add(food string)
}

// 定义乌鸦结构体
type Crow struct {
	name string
}

// -------- 下面开始实现Bird接口 ------
// -------- 实现鸟类接口的所有方法，就代表实现了接口 ------
func (c Crow) fly() {
	fmt.Printf("我是 %s,我会飞....\n", c.name)
}

func (c Crow) eat(food string) {
	fmt.Printf("我是 %s,我喜欢吃 %s \n", c.name, food)
}

// 添加接口
// -------- 实现鸟类接口的所有方法，就代表实现了接口,才是接口类型
func (c Crow) add(food string) {
	fmt.Printf("我是 %s,我喜欢吃 %s \n", c.name, food)
}

// Crow 必须实现所有方法,才能是指定类型
func addBirder(b Birder) {
	fmt.Printf("类型是:%T\n", b)
}

// 定义一个飞行器接口
type Flying interface {
	getName() string
}

// 定义小鸟结构体
type bird struct {
	name string
}

// bird实现接口
func (b bird) getName() string {
	return b.name
}

// 定义飞机结构体
type aircraft struct {
	name string
}

// aircraft实现接口
func (a aircraft) getName() string {
	return a.name
}

// 定义ufo结构体
type ufo struct {
	name string
}

// ufo实现接口
func (u ufo) getName() string {
	return u.name
}

// 写一个函数，接收不同类型的结构体，并打印不同其方法。接口作为参数
func print(flyList []Flying) {
	for _, v := range flyList {
		fmt.Printf("我是%s,我会飞.....\n", v.getName())
	}
}

func main() {

	fmt.Println("\n 2.2 实现接口 ===================================================")
	{
		// Crow 必须实现 所以方法,才表示是接口了了类型
		crow := Crow{"乌鸦"}

		addBirder(crow)
	}

	/*
		Go没有implements或extends关键字,类型都是隐式实现接口的。任何定义了接口中所有方法的类型都被称为隐式地实现了该接口。
	*/
	fmt.Println("\n 2.2 实现接口 ===================================================")
	{
		crow := Crow{"乌鸦"}
		//crow.fly()
		crow.eat("谷子")
	}
	/** 输出
	  我是 乌鸦,我会飞....
	  我是 乌鸦,我喜欢吃 谷子
	*/

	fmt.Println("\n 3. 模拟多态 ===================================================")
	/*
		3. 模拟多态
		3.1 什么是多态?

		如果有几个相似而不完全相同的对象，有时人们要求在向它们发出同一个消息时，它们的反应各不相同，
		分别执行不同的操作，这种情况就是多态现象。Go语言中的多态性是在接口的帮助下实现的——定义接口类型，创建实现该接口的结构体对象。
	*/
	{
		// 定义一个接口切片
		flyList := make([]Flying, 0, 3)
		bird := bird{"小鸟"}
		aircraft := aircraft{"飞机"}
		ufo := ufo{"UFO"}
		flyList = append(flyList, bird, aircraft, ufo)
		fmt.Printf("len: %d cap:%d val: %v \n", len(flyList), cap(flyList), flyList)
		// 调用函数
		print(flyList)
	}
	/** 输出
	  len: 3 cap:3 val: [{小鸟} {飞机} {UFO}]
	  我是小鸟,我会飞.....
	  我是飞机,我会飞.....
	  我是UFO,我会飞.....
	*/

	fmt.Println("\n 4.3 从空接口中取值 ===================================================")

	//{
	//	// 定义一个空接口
	//	type I interface {
	//	}
	//
	//	// 声明变量num
	//	num := 10
	//	// 把变量num存到空接口中
	//	var i I = num
	//	fmt.Printf("输出变量i: %v \n", i)
	//	// 从空接口中取出值，赋值给新的变量
	//	var c int = i // (!!! 这里会报错)
	//	fmt.Printf("输出变量c: %v \n", c)
	//}
	/** 输出
	  ./web.go:16:6: cannot use i (type I) as type int in assignment: need type assertion
	*/

	{

		// 定义一个空接口
		type I interface {
		}

		// 声明变量num
		num := 10
		// 把变量num存到空接口中
		var i I = num
		fmt.Printf("输出变量i: %v \n", i)
		// 从空接口中取出值，赋值给新的变量
		var c int = i.(int)
		fmt.Printf("输出变量c: %v \n", c)
	}
	/**
	  输出变量i: 10
	  输出变量c: 10
	*/

}
