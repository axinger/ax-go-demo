package main

import (
	"fmt"
	"runtime"
)

var ch8 = make(chan int, 6)

func mm1() {
	for i := 0; i < 10; i++ {
		ch8 <- 8 * i
	}
	//close(ch8)
}
func mm2() {
	for {
		for data:=range ch8{
			//fmt.Print(data,"\t")
			fmt.Println("Data = ",data)
		}
	}
}
func main() {
	go mm1()
	go mm2()
	for{
		runtime.GC()
	}
}