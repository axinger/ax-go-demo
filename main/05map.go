package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 引用类型必须初始化

	//var map1 map[string]int

	var map1 = make(map[interface{}]interface{})

	map1["name"] = "jim"
	map1["age"] = 12
	map1[10] = "10"


	fmt.Println("map1 len= ", len(map1))
	fmt.Println("map1 = ", map1)
	// 约定用ok
	v, ok := map1["age"]
	if ok {
		fmt.Println("取值 = ", v)
	} else {
		fmt.Println("不存在")
	}

	for s, i := range map1 {
		fmt.Println(s, i)
	}

	// 随机数,需要买种子
	rand.Seed(time.Now().Unix())
	i := rand.Intn(100)
	fmt.Println("随机数 = ", i)
}
