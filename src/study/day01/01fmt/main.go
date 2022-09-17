package main

import (
	"errors"
	"fmt"
)

func main() {

	// Sprint 返回拼接字串
	s1 := fmt.Sprint("沙河小王子")
	name := "沙河小王子"
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("沙河小王子")
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	e := errors.New("原始错误e")
	w := fmt.Errorf("wrap了一个错误%w", e)
	fmt.Println("e=", e)
	fmt.Println("w=", w)

	var name2 interface{} = nil

	fmt.Println("name2 = ", name2)

	var name3 string

	//name3 = nil

	fmt.Println("name3 = ", len(name3))
}
