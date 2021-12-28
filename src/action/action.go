package action

import "fmt"

var x int8 = int8(10)

const pi = 3.14
const (
	pi2 = pi
)

func init() {
	fmt.Println("action ==== init",x,pi)
}
func Test(s string)  {
	fmt.Println("包 = ",s,x,pi)
}
