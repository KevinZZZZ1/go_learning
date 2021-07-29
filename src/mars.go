package main

import "fmt"

func main() {

	// 无论是数组赋值给新的变量还是将它传递给函数，都会产生一个新的完整数组副本
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	// 数组赋值给变量
	planetsCopy := planets

	planets[2] = "whoops"

	fmt.Println(planets)
	fmt.Println(planetsCopy)

	// 数组作为函数的参数
	terraform(planets)
	fmt.Println(planets)

}

// 这里指定了长度为8的数组，如果传入参数变成其他长度的数组会在编译期间报错，这是因为go中的数组长度也是数组类型的一部分
func terraform(planets [8]string) {
	for i, planet := range planets {
		planets[i] = "New " + planet
	}
}
