package main

import "fmt"

type PolicyType int32

const(
	x PolicyType = iota  // x == 0
	y PolicyType = iota  // y == 1
	z PolicyType = iota  // z == 2
	w  // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说 w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
)
const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0


func main() {

	fmt.Println("枚举值:",x)
	fmt.Println("枚举值:",w)

	fmt.Println("每遇到一个const关键字，iota就会重置:",v)
	fmt.Println("每遇到一个const关键字，iota就会重置:",v)
}
