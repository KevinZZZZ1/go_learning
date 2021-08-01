package main

import (
	"fmt"
	"time"
)

func main() {

	// go中的独立认为叫做goroutine，类似于Java中的线程
	// goroutine的创建效率非常高

	// 在任何方法和函数调用之前使用go关键字就能启动goroutine
	// 向goroutine创建参数和向函数传递参数一样，参数都是按值传递的
	for i := 0; i < 5; i++ {
		go sleepyGopher(i)
	}

	time.Sleep(4 * time.Second)
}

func sleepyGopher(id int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...snore...", id)
}
