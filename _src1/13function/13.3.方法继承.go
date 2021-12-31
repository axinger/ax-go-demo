package main

import "fmt"

// 定义一个人类结构体
type People struct {
	name, position string
	age            int
}

// 定义一个学生结构体
type Student struct {
	People // 继承了People的属性
}
type Teacher struct {
	People // 继承了People的属性
}

// 定义一个方法,这里不需要修改值,所以不用指针
func (p People) say() {
	fmt.Printf("People方法 %s  %d岁 从事: %s \n", p.name, p.age, p.position)
}

// 学生(子类)重写People(父类)的say方法
func (s Student) say() {
	/// 调用父类方法
	s.People.say()
	fmt.Printf("Student方法,名字叫: %s 今年: %d岁 \n", s.name, s.age)
}

func main() {
	{
		student := Student{People{"张三", "学生", 15}}
		teacher := Teacher{People{"李杨", "老师", 35}}

		s := Student{}
		s.name = "王五"
		// 调用方法(继承父类)
		student.say()
		teacher.say()
	}
	/** 输出
	  我叫 张三  15岁 从事: 学生
	  我叫 李杨  35岁 从事: 老师
	*/

	fmt.Println("\n 4 .方法重写 ===================================================")

	{
		student := Student{People{"张三", "学生", 15}}

		// 调用方法(重写父类方法)
		student.say()

		teacher := Teacher{People{"李杨", "老师", 35}}
		// 调用方法(继承父类)
		teacher.say()
	}
	/** 输出
	  我是一名学生,名字叫: 张三 今年: 15岁
	  我叫 李杨  35岁 从事: 老师
	*/
}
