package main

import (
	"errors"
	"fmt"
)

/*
错误是指程序中出现不正常的情况，从而导致程序无法正常运行。
Go语言中没有try...catch来捕获错误，而是通过defer+recover+panic模式来实现捕捉错误信息。

*/

func printError(err error) {
	if err != nil {
		fmt.Printf("err==> %v | err.Error() ==> %v | 类型==> %T \n", err, err.Error(), err)
	}
}

// 创建error对象
func createError(way int) error {
	if way == 1 {
		// 方式一: 使用errors包下的New()
		return errors.New("方式一: 使用errors包下的New() ")
	} else if way == 2 {
		// 方式二: 使用fmt包下的Errorf()
		return fmt.Errorf("方式二: 使用fmt包下的Errorf(...) ---> ")
	}
	return nil
}

/**输出
err==> 方式一: 使用errors包下的New()  | err.Error() ==> 方式一: 使用errors包下的New()  | 类型==> *errors.errorString
err==> 方式二: 使用fmt包下的Errorf(...) --->  | err.Error() ==> 方式二: 使用fmt包下的Errorf(...) --->  | 类型==> *errors.errorString
*/

func main() {

	/*

		2.2 函数返回错误

		对于大多数函数，如果要返回错误，大致上都可以定义为如下模式，⚠️ 必须将error作为多种返回值中的最后一个。

		// 函数返回错误
		func Demo(参数列表...)(x T， err error){
		  // 函数体
		}

	*/

	// 方式一: 使用errors包下的New()
	err := createError(1)
	printError(err)
	//  方式二: 使用fmt包下的Errorf()
	err2 := createError(2)
	printError(err2)
}
