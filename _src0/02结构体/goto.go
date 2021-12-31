package main

import "fmt"

func main() {
	for i := 1; i < 5 ; i++  {
		if i == 3 {
			// 当i=3时，重新循环
			goto RESET
		}
		fmt.Printf("i = %d \n" , i)
	}
	fmt.Println("这里将不会打印出来.....")
RESET:
	fmt.Println("程序从这里开始执行.....")
}
/** 输出
i = 1
i = 2
程序从这里开始执行.....
*/
