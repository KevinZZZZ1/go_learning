package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 使用select来处理多个通道，如果不使用make函数初始化的话，通道变量的值就是nil
	// 对nil通道发送和接收不会panic，但是会永久阻塞
	// 对nil通道执行close函数，会引起panic
	c := make(chan int)

	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}

	timeout := time.After(2 * time.Second)

	for i := 0; i < 5; i++ {
		select {
		case gopherId := <-c:
			fmt.Println("gopher ", gopherId, " has finished sleeping")
		case <-timeout:
			fmt.Println("my patience ran out")
			return
		}
	}
}

func sleepyGopher(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	c <- id
}
