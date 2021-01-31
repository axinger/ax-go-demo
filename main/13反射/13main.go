package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
 map1 := make(map[string]interface{})

 map1["name"]= "jim"
 map1["age"] = 12

	//jsonStr, err := json.Marshal(map1)
	//if err!=nil {
	//	return
	//}
	jsonStr :=`{"name":"jim","age":12}`
	fmt.Println("json = ",string(jsonStr))
	var p person
	json.Unmarshal([]byte(jsonStr),&p)
	fmt.Println("person = ",p)
}
