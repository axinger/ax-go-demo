// 包名 编译成可执行文件
package main

import "fmt"

//批量声明
var (
	name string
	age  int
	isOk bool
)

func main() {

	// = 使用必须使用先var声明例如：
	var a int
	a = 100
	//或
	var b = 100
	//或
	var c int = 100

	// := 是声明并赋值，并且系统自动推断类型，不需要var关键字  ,只能在函数中使用
	d := 100

	name := "jim"
	fmt.Println(a, b, c, d)
	fmt.Println("name = ", name)

	var age = 1
	fmt.Println(age)

	isOk = true

	fmt.Print("Print 在终端中输出要打印的内容 \n")
	fmt.Printf("格式化,go一般用查看类型等操作,但不换行: %s \n", name)

	fmt.Println("直接打印,不能格式化,可以拼接,带换行", "aa")

	// _ 匿名变量名
	/// 调用函数  :=
	// _,res2 := test()
	/// 含有 var
	var _, res2 = test()
	fmt.Println(res2)

}
func test() (string, int) {
	return "jim", 10
}
