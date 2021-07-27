package main

import "fmt"

// 匿名函数就是没有名字的函数，在GO里也可以称为函数字面值

var f func() = func() {
	fmt.Println("Dress up for the masquerade")
}

func main() {

	f()

	// 匿名函数
	d := func(message string) {
		fmt.Println(message)
	}
	d("Go to the party")

	// 声明完匿名函数后立即执行
	func() {
		fmt.Println("Functions anonymous")
	}()

}
