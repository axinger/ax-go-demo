package main

import "fmt"

// 定义一个结构体
type student struct {
	name string
	age  int
}

// 定义一个方法(接收器为Student的指针)
// 接收器变量：接收器变量在命名时，官方建议使用接收器类型的第一个小写字母
func (s *student) updateName(newName string) {
	s.name = newName
}

// 定义一个方法(接收器为Student)
// **若方法的接受者不是指针，实际只是获取了一个拷贝，而不能真正改变接受者中原来的数据。
// ⚠️ 这个需要根据实际情况 决定是否使用 指针
func (s student) updateAge(newAge int) int {
	s.age = newAge
	fmt.Printf("修改结构体s的age不会成功 -> %v \n", s)

	return s.age
}

func main() {

	func() {
		fmt.Println("方法内定义匿名方法,并调用")

		func() {
			fmt.Println("方法内定义匿名方法,并调用")
		}()

	}()

	/*
			Go没有沿袭传统面向对象编程中的诸多概念，也没有提供类(class)，
			但是它提供了结构体(struct)，方法(method)可以在结构体上添加。与类相似，结构体提供了捆绑数据和方法的行为。



			结构体内部无法定义 方法 ,只能在结构体外部定义,通过 func( 结构体名称 )


		func (接收器变量 接收器类型) 方法名(参数列表) (返回参数) {
		     函数体
		}


	*/
	fmt.Println("\n 13方法 ===================================================")

	{
		// 初始化结构体
		s := student{"张三", 20}
		fmt.Printf("结构体初始化s -> %v \n", s)

		// 通过方法修改名称
		s.updateName("张三新名")
		fmt.Printf("调用updateName后 -> %v \n", s)

		// 通过方法修改年龄,无法修改成功,因为 updateAge 方法 接受者不是指针
		age := s.updateAge(22)
		fmt.Printf("调用updateAge后 -> %v ,方法返回值:%v\n", s, age)
	}
	/** 输出:
	  结构体初始化s -> {张三 20}
	  调用updateName后 -> {张三新名 20}
	  修改结构体s的age -> {张三新名 22}
	  调用updateAge后 -> {张三新名 20}
	*/

	/*
		3.方法继承

		方法是可以继承的，如果匿名字段实现了一个方法，那么包含这个匿名字段的struct也能调用该匿名字段中的方法。
	*/
	fmt.Println("\n 3.方法继承 ===================================================")

}
