package main

import "fmt"

type animal interface {
	//move()
	//eat()

	mover
	eater // 接口嵌套
}

type mover interface {
	move()
}

type eater interface {
	eat()
}

type dog struct {
}

func (d *dog) move() {

}
func (d *dog) eat() {

}

func test(a animal) {

	fmt.Printf("animal 类型: %T \n", a)
}

func main() {

	d := dog{}

	test(&d)
	fmt.Print(d)
}
