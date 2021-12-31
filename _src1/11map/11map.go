package main

import "fmt"

func main() {

	fmt.Println("\n 11map ===================================================")

	fmt.Println("\n 2.1 声明初始化 ===================================================")
	{
		// 声明的同时初始化
		var ageMap = map[string]int{
			"张三": 20,
			"李四": 23,
			"王五": 33,
		}
		fmt.Printf("变量ageMap--> 值: %v 类型: %T \n", ageMap, ageMap)

		// 短变量声明初始化
		ageMap2 := map[string]int{"张三": 20, "李四": 23, "王五": 33}
		fmt.Printf("变量ageMap2--> 值: %v 类型: %T", ageMap2, ageMap2)
	}
	/**输出
	  变量ageMap--> 值: map[张三:20 李四:23 王五:33] 类型: map[string]int
	  变量ageMap2--> 值: map[张三:20 李四:23 王五:33] 类型: map[string]int
	*/

	fmt.Println("\n 2.2 使用make ===================================================")
	{
		// 先创建后赋值
		ageMap := make(map[string]int)
		ageMap["张三"] = 20
		ageMap["李四"] = 23
		ageMap["王五"] = 23
		fmt.Printf("变量ageMap-- 值: %v 类型: %T", ageMap, ageMap)
	}
	// 输出: 变量ageMap--> 值: map[张三:20 李四:23 王五:23] 类型: map[string]int

	fmt.Println("\n 3.1 遍历基础map ===================================================")
	{
		// 短变量声明初始化
		ageMap := map[string]int{"张三": 20, "李四": 23, "王五": 33}
		// 遍历
		for k, v := range ageMap {
			fmt.Printf("k: %v ,  v:%v \n", k, v)
		}

		for s := range ageMap {
			fmt.Printf("s>> 值: %v \n", s)
		}
	}
	/**输出
	  姓名: 张三 年龄: 20
	  姓名: 李四 年龄: 23
	  姓名: 王五 年龄: 33
	*/

	fmt.Println("\n 4.1 判断key是否存在 ===================================================")
	{
		// 声明map
		fruitMap := map[string]float32{"香蕉": 3.22, "苹果": 1.88, "葡萄": 2.49}

		// 存在: ok=true  存在: ok=false
		price, ok := fruitMap["香蕉"]
		if ok {
			fmt.Printf("香蕉存在！价格: %.2f \n", price)
		}
		price2, ok2 := fruitMap["樱桃"]
		if ok2 {
			fmt.Printf("樱桃存在！价格: %.2f \n", price2)
		} else {
			fmt.Printf("樱桃不存在! \n")
		}
		// 简写
		if price, ok := fruitMap["苹果"]; ok {
			fmt.Printf("苹果存在！价格: %.2f \n", price)
		}

		// 简写
		if _, ok := fruitMap["苹果1"]; !ok {
			fmt.Printf("苹果1不存在")
		}

	}
	/**输出
	  香蕉存在！价格: 3.22
	  樱桃不存在!
	  苹果存在！价格: 1.88
	*/

	fmt.Println("\n 4.2 删除 判断key是否存在 ===================================================")
	/*
		delete(map, key) 函数用于删除集合的某个元素，参数为map和其对应的key。删除函数不返回任何值。
	*/
	{
		// 声明map
		fruitMap := map[string]float32{"香蕉": 3.22, "苹果": 1.88, "葡萄": 2.49, "梨": 4.13}
		fmt.Printf("删除前-->fruitMap = %v \n", fruitMap)
		// 删除苹果
		delete(fruitMap, "苹果")
		fmt.Printf("删除后-->fruitMap = %v \n", fruitMap)
		// 清空map
		fruitMap = map[string]float32{}
		fmt.Printf("清空后-->fruitMap = %v \n", fruitMap)
	}
	/**输出
	  删除前-->fruitMap = map[梨:4.13 苹果:1.88 葡萄:2.49 香蕉:3.22]
	  删除后-->fruitMap = map[梨:4.13 葡萄:2.49 香蕉:3.22]
	  清空后-->fruitMap = map[]
	*/

	/*
	   5.map是引用类型

	   map与切片相似，都是引用类型。将一个map赋值给一个新的变量时，它们指向同一块内存（底层数据结构）。因此，修改两个变量的内容都能够引起它们所指向的数据发生变化。
	*/

	{
		// 声明map
		fruitMap := map[string]float32{"香蕉": 3.22, "苹果": 1.88, "葡萄": 2.49, "梨": 4.13}
		fmt.Printf("调用函数前-->fruitMap = %v \n", fruitMap)
		// 调用函数
		testMap(fruitMap)
		fmt.Printf("调用函数后-->fruitMap = %v \n", fruitMap)
	}

	/**输出
	  调用函数前-->fruitMap = map[梨:4.13 苹果:1.88 葡萄:2.49 香蕉:3.22]
	  调用函数后-->fruitMap = map[梨:4.13 苹果:1.88 葡萄:2.49 香蕉:5.99]
	*/

}

// map是引用类型，函数改变他的值外层变量也会变
func testMap(fruitMap map[string]float32) {
	fruitMap["香蕉"] = 5.99
}
