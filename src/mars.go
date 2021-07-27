package main

import "fmt"

// 匿名函数就是没有名字的函数，在GO里也可以称为函数字面值

var f func() = func() {
	fmt.Println("Dress up for the masquerade")
}

func main() {

	f()

}
