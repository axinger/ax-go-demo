package main

import (
	"encoding/json"
	"fmt"
)

type myInt int     // 自定义别名
type yourInt = int // 类型别名
/// go 标识符 首字母大写,就表示对外部可见,共有的public的类型
// 不能给别的包里面的类添加方法,只能自己包
// go 没有继承

// 结构体
type Person struct {
	name  string
	age   int
	hobby []string
	//int // 结构体匿名字段 .属性名取值 ,不能重复
}

type Student struct {
	name  string
	age   int
	hobby []string
	dog   Dog
	Book  // 匿名嵌套结构体
}

// 模拟动物行为的接口
type IAnimal interface {
	Eat() // 描述吃的行为
}

// 动物,模拟继承,
type Animal struct {
	class string
}

type Dog struct {
	// 属性大写,才能json 序列化, 后面对应json的字段
	Name string `json:"name"`
	name    string
	dogName string
	Animal  // 模拟继承
}

type Book struct {
	name     string
	bookName string
}

//构造函数 new开头
// 当结构体大的时候,使用结构体指针
func newStudent(name string, age int) *Student {
	return &Student{
		name: name,
		age:  age,
	}
}

// 参数是 值拷贝
func test1(p Person) {
	p.name = "name$test1"
}

// 指针类型
func test2(p *Person) {
	//(*p).name = "name$test2"
	p.name = "name$test2" // 语法糖格式
}

// 方法和接收者
// 方法作用特定类型的函数 函数名前面写
// 接收者 具体调用的类型 (p Person),字母小写
func (p Person) eat() {
	fmt.Println("吃饭...", p.name)
}

// 指针类型
func (p *Person) eat2() {
	p.name = "name_eat2"
	fmt.Println("吃饭...", p.name)
}

func (i myInt) myInt() {
	fmt.Println("int 添加方法...", i)
}

//
//// 接收者 具体调用的类型 (p Person),字母小写
//func ([p Person,s Student]])eat2()  {
//	fmt.Println("吃饭...",p.name)
//}

func main() {
	var a yourInt = 100
	fmt.Printf("yourInt = %T \n", a)

	// 声明变量并初始化,有key
	var p1 = Person{
		name:  "tom",
		age:   1,
		hobby: []string{"1"},
	}

	//p1.name = "jim"
	//p1.age = 10
	p1.hobby = append(p1.hobby, "3")
	fmt.Println("p1 = ", p1)

	p1.eat()

	st1 := Student{
		name: "jim",
		dog: Dog{
			name: "tom",
		},
		// 初始化用类型名称
		Book: Book{
			bookName: "book",
		},
	}
	fmt.Println("st1.name = ", st1.name)
	fmt.Println("st1.dog.name = ", st1.dog.name)

	fmt.Println("st1.bookName = ", st1.bookName)

	/// 指针
	var p3 = new(Person)
	p3.name = "name,p3"
	fmt.Println("p3 new构造 = ", p3)

	// 指针
	var p4 = &Person{}
	p4.name = "name,p4"
	fmt.Println("p4 = ", p4)

	// 没有key,
	var p5 = Person{
		"tom",
		1,
		[]string{"1"},
	}
	fmt.Println("p5 = ", p5)

	//var p2 =  p1
	//p1.name = "tom2"
	test1(p1)
	fmt.Println("p1修改 = ", p1)
	//fmt.Println("p2 = ",p2)

	test2(&p1)
	fmt.Println("p1修改指针 = ", p1)

	//var s = struct {
	//	name string
	//}{}
	// 匿名结构体
	var s struct {
		name string
	}
	s.name = "jim"
	fmt.Println("匿名结构体 = ", s)

	s1 := newStudent("jim", 10)
	fmt.Println("p10 = ", s1)
	s2 := s1
	s2.name = "jim_11"
	fmt.Println("p10 修改 = ", s1)
	// 内置int添加方法
	i1 := myInt(10)
	i1.myInt()

	d3 := Dog{
		name: "tom",
		Name: "Tom",
		// 模拟继承
		Animal: Animal{
			class: "犬科",
		},
	}
	fmt.Println("d3.class = ", d3.class)

	js1, err := json.Marshal(d3)


	if err != nil { //  表示有错误
		fmt.Println("对象转json错误 = ",err)
		return
	}
	//fmt.Println("对象转err = ",err)
	fmt.Println("对象转json = ",string(js1))

	jsStr := string(js1)
	d4 := Dog{}
	error2 := json.Unmarshal([]byte(jsStr), &d4)
	if error2 != nil { //  表示有错误
		fmt.Println("json转对象错误 = ",err)
		return
	}
	fmt.Println("json转对象 name = ",d4.name,",","Name =",d4.Name)
}
