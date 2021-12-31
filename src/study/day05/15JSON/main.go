package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	// json 解析 需要大写
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 序列化
	p1 := person{
		Name: "jim",
		Age:  10,
	}
	fmt.Println("person = ", p1)

	json1, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("对象转json#v 含有斜线格式化 = %#v \n", string(json1))
	fmt.Printf("对象转json = %v \n", string(json1))

	// 反序列化
	jsonStr2 := `{"name":"jim","age":12}`
	var p2 person
	json.Unmarshal([]byte(jsonStr2), &p2)
	fmt.Printf("json转对象 = %#v \n", p2)

}
