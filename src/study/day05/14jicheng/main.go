package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) eat() {
	fmt.Println("person = eat", "name = ", p.name)
}

type student struct {
	person
}

func (s student) eat() {
	//s.person.eat() // 重写父类
	fmt.Println("student = eat", "name = ", s.name)
}

func main() {

	s1 := student{person{
		name: "jim",
		age:  0,
	}}

	fmt.Println("s =", s1.name)

	//s1.person.eat()
	s1.eat()
}
