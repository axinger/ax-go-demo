package main

import "fmt"

/*

5.1 转换语法

// 方式一
instance,ok := 接口对象.(实际类型)
// 方式二
接口对象.(实际类型)

*/

func main() {

	// 声明变量
	var a I
	// 保存整型
	a = 10
	printType(a)
	printType2(a)
	// 保存字符串
	a = "hello word"
	printType(a)
	printType2(a)
	// 保存数组
	a = [3]float32{1.0, 2.0, 3.0}
	printType(a)
	printType2(a)
	// 保存切片
	a = []string{"您", "好"}
	printType(a)
	printType2(a)
	// 保存Map
	a = map[string]string{
		"张三": "男",
		"小丽": "女",
	}
	printType(a)
	printType2(a)
	// 保存结构体
	a = people{"刘山", 32}
	printType(a)
	printType2(a)

}

// 定义一个空接口
type I interface {
}

// 定义结构体
type people struct {
	name string
	age  int
}

// 方式一
func printType(i I) {
	if t, ok := i.(int); ok {
		echo(t)
	} else if t, ok := i.(string); ok {
		echo(t)
	} else if t, ok := i.(map[string]string); ok {
		echo(t)
	} else if t, ok := i.([]int); ok {
		echo(t)
	} else if t, ok := i.([3]string); ok {
		echo(t)
	} else if t, ok := i.(people); ok {
		echo(t)
	}
}

// 方式二
func printType2(i I) {
	switch i.(type) {
	case int:
		echo2(i)
	case string:
		echo2(i)
	case map[string]string:
		echo2(i)
	case []int:
		echo2(i)
	case [3]string:
		echo2(i)
	case people:
		echo2(i)
	}
}
func echo(i interface{}) {
	fmt.Printf("方式一 ---> 变量i类型: %T 值: %v \n", i, i)
}
func echo2(i interface{}) {
	fmt.Printf("方式二 ---> 变量i类型: %T 值: %v \n", i, i)
}
