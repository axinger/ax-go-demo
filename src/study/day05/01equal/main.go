package main

import (
	"fmt"
	"reflect"
)

func main() {

	s1 := map[string]interface{}{
		"name": "jim",
	}

	s2 := map[string]interface{}{
		"name": "jim",
	}

	s3 := map[string]interface{}{
		"name": "tom",
	}

	fmt.Println("引用类型判断值是否一致:", reflect.DeepEqual(s1, s2))
	fmt.Println("引用类型判断值是否一致:", reflect.DeepEqual(s1, s3))
}
