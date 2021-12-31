package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type person struct {
	Name string `json:"name" json2:"name2"`
	Age  int    `json:"age"`
}

func main() {
	map1 := make(map[string]interface{})

	map1["name"] = "jim"
	map1["age"] = 12

	//jsonStr, err := json.Marshal(map1)
	//if err!=nil {
	//	return
	//}
	jsonStr := `{"name":"jim","age":12}`
	fmt.Println("json = ", string(jsonStr))
	var p person
	json.Unmarshal([]byte(jsonStr), &p)
	fmt.Println("person = ", p)

	// 获得类型 子类型
	/// .Kind() 根类
	class := reflect.TypeOf(person{})

	for i:=0;i<class.NumField();i++ {
		field := class.Field(i)
		fmt.Println(" TypeOf NumField = ",field.Name,field.Index,field.Type,field.Tag.Get("json"),field.Tag.Get("json2"))
	}


	fmt.Println("类型 = ", class, class.Name(), class.Kind())

	/// 获得 值 这里传指针
	value := reflect.ValueOf(&p)

	for i:=0;i<reflect.ValueOf(p).NumField();i++ {
		field := class.Field(i)
		fmt.Println(" value NumField = ",field.Name,field.Index,field.Type,field.Tag.Get("json"),field.PkgPath)
		//fmt.Println(" value NumField = ",field)
	}




	fmt.Println("反射值 = ", value)

	elName := value.Elem().FieldByName("Name")

	if elName.Kind() == reflect.String {
		fmt.Println("判断属性类型 string")
		elName.SetString("tom")
	}

	elAge := value.Elem().FieldByName("Age")
	if elAge.Kind() == reflect.Int {
		elAge.SetInt(10)
	}

	elH := value.Elem().FieldByName("h")
	fmt.Println("反射修改值 elH = ", elH.IsValid())
	if elH.Kind() == reflect.Int {
		elH.SetInt(10)
	}


	fmt.Println("反射修改值 = ", p)


	// 结构体反射


}
