package main

import (
	"fmt"
	"time"
)

/*
2. 创建Channel
2.1 语法

channel是引用类型，需要使用make()进行创建。

// 声明方式1
var cha1 chan 数据类型
cha1 = make(chan 数据类型)

// 声明方式2
cha1 := make(chan 数据类型)

// 读写
intChan := make(chan int)

// 只读
ch <- chan T
// 只写
ch chan <- T


*/

type People struct{}

func main() {

	{
		// 创建一个整数型chan
		intChan := make(chan int)
		fmt.Printf("intChan类型: %T 值: %v \n", intChan, intChan)
		// 创建一个空接口chan,可以存放任意类型数据
		interfaceChan := make(chan interface{})
		fmt.Printf("interfaceChan类型: %T 值: %v \n", interfaceChan, interfaceChan)
		// 创建一个指针chan
		peopleChan := make(chan *People)
		fmt.Printf("peopleChan类型: %T 值: %v \n", peopleChan, peopleChan)

		// 3.发送数据
		/*
				channel变量名 <- 值

			⚠️  如果Goroutine在一个channel上发送数据，其他的Goroutine应该接收得到数据；如果没有接收，那么程序将在运行时出现死锁。
		*/

		// 写入,必须先写入,才能接收
		// 一个channel是一条通信管道，它可以让一个协程通过它给另一个协程发送数据
		go func() {
			intChan <- 5
			fmt.Printf("写入数据: %T 值: %v ,地址:%p\n", intChan, intChan, intChan)
		}()

		//go func() {
		//	intChan <- 6
		//	fmt.Printf("写入数据: %T 值: %v ,地址:%p\n", intChan, intChan, intChan)
		//}()
		// 接收值
		a := <-intChan
		fmt.Printf("接收数据: %v \n", a)
	}
	/** 输出
	  intChan类型: chan int 值: 0xc000052060
	  interfaceChan类型: chan interface {} 值: 0xc0000520c0
	  peopleChan类型: chan *main.People 值: 0xc000052120
	*/

	//4.普通接收
	/*
	   // 方式一: ch 指的是通道变量
	   data := <- ch
	   //方式二: data 表示接收到的数据。未接收到数据时，data为channel类型的零值，ok（布尔类型）表示是否接收到数据
	   data,ok := <- ch
	*/

	{
		// 创建一个整数型chan
		intChan := make(chan int)
		// 写入数据(协程写入)
		go func(cha2 chan int) {
			// 写入
			for i := 1; i < 5; i++ {
				cha2 <- i
				fmt.Printf("写入数据 ->  %v \n", i)
			}
			// 关闭通道
			close(cha2)
			//close(intChan)
		}(intChan)

		// 方式二: data,ok := <- ch
		//for {
		//	// 接收数据
		//	out,ok := <-intChan
		//	// 判断通道是否关闭,如果通道关闭，则ok为false
		//	if !ok {
		//		fmt.Println("通道已关闭")
		//		break
		//	}
		//	fmt.Printf("接收数据 ==> %v \n", out)
		//}
		//fmt.Println("程序运行结束!")

		// 使用 for...range接收
		for data := range intChan {
			fmt.Printf("使用 for...range接收 接收数据 ==> %v \n", data)
		}
		fmt.Println("程序运行结束!")

	}
	/** 输出:
	  写入数据 ->  1
	  接收数据 ==> 1
	  接收数据 ==> 2
	  写入数据 ->  2
	  写入数据 ->  3
	  接收数据 ==> 3
	  接收数据 ==> 4
	  写入数据 ->  4
	  通道已关闭
	  程序运行结束!
	*/

	{
		// 创建一个整数型chan
		intChan := make(chan int)
		// 创建一个用于阻塞的chan
		boolChan := make(chan bool)

		// 创建一个写入数据的协程
		go func(cha chan int) {
			// 写入
			intChan <- 50
			fmt.Println("写入数据50")
			// 关闭通道
			close(intChan)
		}(intChan)

		// 创建一个读取数据的协程
		go func(intChan chan int, boolChan chan bool) {
			data, ok := <-intChan
			if ok {
				fmt.Printf("读取到数据 ->  %v \n", data)
				// 读取到数据后，给boolChan写入值
				boolChan <- true
				// 关闭用于的阻塞的chan
				close(boolChan)
			}
		}(intChan, boolChan)
		// 忽略接收，达到阻塞的效果。(如果不阻塞，则会直接输出: 程序运行结束!,不会等待协程执行)
		<-boolChan
		fmt.Println("程序运行结束!")
	}
	/** 输出
	  写入数据50
	  读取到数据 ->  50
	  程序运行结束!
	*/

	{
		// 创建一个整数型chan
		intChan := make(chan int)
		// 创建一个写入channel的协程
		go func(intChan chan int) {
			intChan <- 10
			intChan <- 20
			// 关闭通道
			close(intChan)
		}(intChan)

		// 读取数据
		a := <-intChan
		fmt.Printf("接收数据: %v \n", a)
		//b := <- intChan
		//fmt.Printf("接收数据: %v \n",b)
		//
		//// 此时的Chan已经关闭，而且里面的数据也都已经取完
		//c := <- intChan
		//fmt.Printf("接收数据: %v \n",c)
		//fmt.Println("程序运行结束!")
	}
	/** 输出
	  接收数据: 10
	  接收数据: 20
	  接收数据: 0
	  程序运行结束!
	*/

	/*
		默认创建的都是非缓冲channel，读写都是即时阻塞。缓冲channel自带一块缓冲区，
		可以暂时存储数据，如果缓冲区满了，就会发生阻塞。
		缓冲通道在发送时无需等待接收方接收即可完成发送过程，并且不会发生阻塞，只有当缓冲区满时才会发生阻塞。
		同理，如果缓冲通道中有数据，接收时将不会发生阻塞，直到通道中没有数据可读时，通道将会再度阻塞。
	*/
	{
		fmt.Printf("开始时间: %v \n", time.Now().Unix())
		// 创建一个缓冲区为2的整数型chan
		intChan2 := make(chan int, 2)
		// 不会发生阻塞，因为缓冲区未满
		intChan2 <- 100
		//intChan2 <- 100

		// 超过2个缓存,没有结束者,会阻塞奔溃
		//num2 := <-intChan2
		//fmt.Println("num2 = ",num2)
		//intChan2 <- 100

		fmt.Printf("结束时间: %v \n", time.Now().Unix())
		fmt.Printf("intChan2  类型: %T 缓冲大小: %v \n", intChan2, cap(intChan2))
		fmt.Println("程序运行结束!")
	}
	/**输出:
	  开始时间: 1607496281
	  结束时间: 1607496281
	  intChan2  类型: chan int 缓冲大小: 2
	  程序运行结束!
	*/

	/*
		9 只读,只写
	*/

	{
		// 创建一个整数型chan
		//intChan := make(chan int)
		//go writeChan(intChan)
		//go readChan(intChan)
		//time.Sleep(50 * time.Millisecond)
		//fmt.Println("运行结束")

	}

	/*
		10.计时器与channel

		计时器类型表示单个事件。当计时器过期时，当前时间将被发送到c上（c是一个只读channel ＜-chan time.Time，该channel中放入的是Timer结构体），
		除非计时器是After()创建的。计时器必须使用NewTimer()或After()创建。

	*/

	{
		// 创建一个计时器,返回的是chan

		/*
			 返回 只读通道
			func After(d Duration) <-chan Time {
				return NewTimer(d).C
			}

			// 只读
			ch <- chan T
			// 只写
			ch chan <- T


		*/
		ch := time.After(5 * time.Second)
		fmt.Printf("开始时间 %v \n", time.Now())
		fmt.Println("开始阻塞=============")
		// 此处会阻塞5秒
		out := <-ch
		fmt.Println("结束阻塞=============")
		fmt.Printf("变量out->  类型: %T 值:%v  \n", out, out)
		fmt.Printf("开始时间 %v \n", time.Now())
	}

	{
		// 创建一个计时器,返回的是chan
		fmt.Printf("开始时间 %v \n", time.Now())
		fmt.Println("开始睡眠阻塞=============")
		// 此处会阻塞5秒
		time.Sleep(5 * time.Second)
		fmt.Println("结束睡眠阻塞=============")
		fmt.Printf("开始时间 %v \n", time.Now())
	}

}
