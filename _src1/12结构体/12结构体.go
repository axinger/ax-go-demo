package main

import "fmt"

// 定义结构体
type student struct {
	name string
	age  int
	like []string
}

func main() {
	fmt.Println("\n12结构体===================================================")
	{
		// 使用var 实例化结构体
		var s student
		fmt.Printf("变量s--> 类型: %T 值: %v \n", s, s)
		// 给属性赋值
		s.name = "张三"
		s.age = 20
		s.like = []string{"打游戏", "看动漫"}
		fmt.Printf("给属性赋值: 变量s--> 类型: %T 值: %v \n", s, s)
	}

	// 方式一: 先实例化结构体，后赋值
	s := student{}
	// 给属性赋值
	s.name = "张三"
	s.age = 20
	s.like = []string{"打游戏", "看动漫"}
	fmt.Printf(" 变量s--> 类型: %T 值: %v \n", s, s)

	// 方式二: 声明时初始化
	s1 := student{
		name: "李四",
		age:  23,
		like: []string{"旅游", "运动"},
	}
	fmt.Printf(" 变量s1--> 类型: %T 值: %v \n", s1, s1)

	// 方式三: 声明时初始化，省略属性
	s2 := student{"王麻子", 30, []string{"睡觉", "吃饭"}}
	fmt.Printf(" 变量s2--> 类型: %T 值: %v \n", s2, s2)

	// 3.3 使用new
	/**
	使用内置函数new()对结构体进行实例化，结构体实例化后形成指针类型的结构体，
	new()内置函数会分配内存。第一个参数是类型，而不是值，返回的值是指向该类型新分配的零值的指针。
	*/
	{
		// 使用new实例化
		s := new(student)
		fmt.Printf(" 变量s--> 类型: %T 值: %v \n", s, s)
		// 给属性赋值
		(*s).name = "包青天"
		(*s).age = 55
		(*s).like = []string{"判案"}
		fmt.Printf(" 变量s--> 类型: %T 值: %v \n", s, s)
		// 语法糖写法(省略*)
		s.name = "包大人"
		s.age = 99
		s.like = []string{"判案", "元芳你怎么看"}
		fmt.Printf(" 变量s--> 类型: %T 值: %v \n", s, s)

	}

	fmt.Println("\n结构体作为参数===================================================")
	// 结构体作为参数,只能是值传值
	{

		s := student{name: "张三", age: 17}
		fmt.Printf(" 变量s--> 值: %v \n", s)
		grownUp(s)
		fmt.Printf("调用函数后,变量s--> 值: %v \n", s)

	}

	{
		s := student{name: "张三", age: 17}
		fmt.Printf(" 变量s--> 值: %v \n", s)
		// 取址
		grownUp2(&s)
		fmt.Printf("结构体指针传值:调用函数后,变量s--> 值: %v \n", s)
	}

	fmt.Println("\n4.匿名结构体===================================================")
	// 4.匿名结构体

	{

		// 声明初始化匿名结构体
		s := struct {
			name, home, phone string
			age               int
		}{
			name:  "张二十",
			phone: "17600111111",
			age:   18,
		}

		s.name = "jim"
		// 打印
		fmt.Printf("匿名结构体 变量 s--> 值: %v  类型: %T \n", s, s)
	}

	fmt.Println("\n5.结构体含有匿名字段===================================================")
	// 5.结构体含有匿名字段
	{
		type people struct {
			name, home string
			int        // 匿名字段
			float32    // 匿名字段
		}

		// 声明初始化匿名结构体
		s := people{name: "张三", home: "北京", int: 18, float32: 1.73}
		fmt.Printf("变量 s--> 值: %v \n", s)
		// 声明初始化匿名结构体(省略属性名)
		s2 := people{"李四", "南京", 22, 1.80}
		fmt.Printf("结构体含有匿名字段 变量 s2--> 值: %v \n", s2)

	}

	fmt.Println("\n6. 结构体嵌套===================================================")

	/*
		6.1 定义

		将一个结构体作为另一个结构体的属性（字段），这种结构就是结构体嵌套。

		结构体嵌套可以模拟面向对象编程中的以下两种关系。

		    聚合关系: 一个类作为另一个类的属性。
		    继承关系: 一个类作为另一个类的子类。子类和父类的关系。

	*/

	fmt.Println("\n6.2 聚合场景===================================================")
	{

		// 定义学习结构体
		type school struct {
			schoolName, schoolAddress string
		}

		// 定义学生结构体
		type student struct {
			name       string
			height     float32
			schoolInfo school
		}
		// 简短声明嵌套结构体
		s := student{"小张", 1.72, school{"北京大学", "北京"}}
		fmt.Printf("变量 s--> 值: %v 类型: %T \n", s, s)
		// 使用var
		var ss student
		ss.name = "小龙"
		ss.height = 1.67
		ss.schoolInfo = school{"南京大学", "南京"}
		fmt.Printf("变量 ss--> 值: %v 类型: %T \n", ss, ss)
		// 使用new
		s2 := new(student)
		s2.name = "小虎"
		s2.height = 1.77
		s2.schoolInfo.schoolName = "武汉大学"
		s2.schoolInfo.schoolAddress = "武汉"
		fmt.Printf("变量 s2--> 值: %v 类型: %T \n", s2, s2)

	}
	fmt.Println("\n6.3 模拟继承===================================================")
	/*
		在结构体中，属于匿名结构体的字段称为提升字段，它们可以被访问，匿名结构体就像是该结构体的父类。
	*/
	{

		// 定义父类结构体
		type people struct {
			name string
			age  int
		}
		type student struct {
			people  // 集成父类结构体,就是直接写父类的类名,作为属性名
			address string
		}

		// 方式1.使用new声明结构体
		var s = new(student)
		// 集成父类成员
		s.name = "张三"
		s.age = 12
		// 自己成员
		s.address = "三年级"
		fmt.Printf("变量s -> %v \n", s)
		// 方式2.使用简短声明
		s2 := student{people{"李四", 13}, "四年级"}
		s2.name = "jim"
		s2.people.name = "tom"
		fmt.Printf("变量s2 -> %v \n", s2)

	}
	fmt.Println("\n6.4 成员冲突===================================================")
	{

		type A struct {
			name string
			age  int
		}
		type B struct {
			name   string
			height float32
		}
		// 在C结构体中嵌套A和B
		type C struct {
			A
			B
		}

		// 定义结构体C
		c := C{}
		// 不冲突的成员赋值
		c.age = 12
		c.height = 1.88
		// 冲突的成员赋值
		c.A.name = "这是A的成员"
		c.B.name = "这是B的成员"
		//c.name 这个只能赋值第一个属性A
		fmt.Printf("变量c -> %v \n", c)

		// 输出: 变量c -> {{这是A的成员 12} {这是B的成员 1.88}}
	}

}

// 传结构体值作为参数
func grownUp(s student) {
	s.age = 80
	s.name = "长大的 " + s.name
}

// 传结构体指针作为参数
func grownUp2(s *student) {
	s.age = 80
	s.name = "长大的 " + s.name
}

// 作为值类型传递
func getStudent(name string, age int, likes []string) student {
	return student{name, age, likes}
}

// 返回指针
func getStudent2(name string, age int, likes []string) *student {
	return &student{name, age, likes}
}
