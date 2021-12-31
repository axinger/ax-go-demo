package main

// User  使用示例
type User struct {
	name    string
	age     int
}


func main() {


//结构体的初始化可以使用new关键词和var关键词，不同的是如果使用new，则返回类型是一个指针，使用var，则是结构体自身。
	user := new(User)
	user.name = "tom"
	user.age = 20


	user2 := User{}
	user2.name = "tom"
	user2.age = 20


	var user3 = User{}
	user3.name = "tom"
	user3.age = 20

}
