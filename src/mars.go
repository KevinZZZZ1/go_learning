package main

import (
	"fmt"
	"time"
)

func main() {
	// 通道(channel)可以用在多个goroutine之间安全的传值
	// 通道可以作为变量，函数参数，结构体字段等
	// 可以使用make函数来创建通道，并且要指定通道传输数据的类型

	// 发送操作的goroutine会等待，直到另一个goroutine尝试对该通道进行接收操作为止
	// 执行接收操作的goroutine将等待直到另一个goroutine尝试向该通道进行发送操作为止
	var c chan int = make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}

	for i := 0; i < 5; i++ {
		gopherId := <-c
		fmt.Println("gopher ", gopherId, " has finished sleeping")
	}
}

func sleepyGopher(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...", id, " snore ...")
	c <- id
}
