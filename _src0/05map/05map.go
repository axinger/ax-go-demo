package main

import (
	"fmt"
	"math/rand"
	"time"
)

// map是引用类型，函数改变他的值外层变量也会变
func testMap(fruitMap map[string]float32){
	fruitMap["香蕉"] = 5.99
}

func main() {

	 map2 := map[string]interface{}{}
	 map2["a"]= 1
	fmt.Printf("map2 = %v \n",map2)

	// 声明map
	fruitMap := map[string]float32{"香蕉": 3.22, "苹果": 1.88, "葡萄": 2.49,"梨":4.13}
	fmt.Printf("调用函数前-->fruitMap = %v \n",fruitMap)
	// 调用函数
	testMap(fruitMap)
	fmt.Printf("调用函数后-->fruitMap = %v \n",fruitMap)


	// 引用类型必须初始化

	//var map1 map[string]int
	_ = make(map[string]int)
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
