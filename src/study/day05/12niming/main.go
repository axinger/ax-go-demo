package main

import "fmt"

type person struct {
	name string
	age  int
	int
}

func main() {

	p1 := person{
		name: "jim",
		age:  10,
		int:  11,
	}
	fmt.Print(p1.int)

}
