package main

import (
	"fmt"
	"time"
)

// 第一步: 定义一个结构体
type MyError struct {
	msg string
	t   time.Time
}

// 第二步: 并实现error接口

/*
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
	Error() string
}

实现接口,直接写接口的方法
*/
func (m MyError) Error() string {
	return fmt.Sprintf("错误信息: %s 发生时间: %v", m.msg, m.t)
}

// 第三步: 定义一个返回error的函数
func login(phone, pwd string) (bool, error) {
	if phone == "17600000111" && pwd == "123456" {
		return true, nil
	}
	err := MyError{"账号密码错误！", time.Now()}
	return false, err
}

func main() {
	// 测试
	res, err := login("17600000111", "123789")
	if err != nil {
		fmt.Printf("登录失败-->  %v T --> %T \n", err.Error(), err)
	} else {
		fmt.Printf("登录成功 -->  %v \n", res)
	}
	res2, err2 := login("17600000111", "123456")
	if err2 != nil {
		fmt.Printf("登录失败-->  %v T --> %T \n", err2.Error(), err2)
	} else {
		fmt.Printf("登录成功 -->  %v \n", res2)
	}
}

/**输出
登录失败-->  错误信息: 账号密码错误！ 发生时间: 2020-11-27 16:09:39.774176 +0800 CST m=+0.000091945 T --> main.MyError
登录成功 -->  true
*/
