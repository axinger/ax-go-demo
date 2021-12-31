package main

import "fmt"

func main() {
	map1 := map[string]interface{}{}

	map1["age"] = 12
	fmt.Println(map1)

	map2 := make(map[string]interface{})
	map2["age"] = 12
	fmt.Println(map2)

	test(1)
	test("a")

}

// 类型断言
func test(a interface{}) {

	switch a.(type) {

	case string:

		fmt.Println(a, "是 string", a.(string))

	case int:
		fmt.Println(a, "是 int", a.(int))
	}
}
