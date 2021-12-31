package main

import "fmt"

func main() {
	err1()
	err2()
	err3()
}

func err1()  {

	fmt.Println("err1.....")
}
func err2()  {
	defer func() {

		err := recover()
		fmt.Println("recover,必须搭配defer使用,少用 恢复现场,后面执行 = ",err)

		fmt.Println("错误,defer,释放资源")
	}()
	panic("出现了严重错误....,后面不会执行")
	fmt.Println("err2.....")
}
func err3()  {
	fmt.Println("err3.....")
}


