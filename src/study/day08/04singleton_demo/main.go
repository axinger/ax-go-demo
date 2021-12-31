package main

import (
	singleton "axinger.xyz/ax_go_demo/study/day08/04singleton"
	"fmt"
)

func main() {

	fmt.Printf("地址%p", singleton.GetInstance())
	fmt.Printf("地址%p", singleton.GetInstance())

}
