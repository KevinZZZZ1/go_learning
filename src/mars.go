package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func main() {
	// 复合字面量的前面可以使用&

	var timmy *person = &person{
		name: "Timmy",
		age:  10,
	}

	// 虽然timmy是一个person类型的指针，按理应该是(*timmy).superpower = "flying"
	// 但是go中有自动解引用的功能，所以下面这种写法等于上面的，即访问struct中的字段的时候，解引用不是必须的操作
	timmy.superpower = "flying"

	fmt.Printf("%+v\n", timmy)

	// 指向数组的指针
	var superpower *[3]string = &[3]string{"flight", "invisibility", "super strength"}
	// 无需进行解引用，直接使用指针
	fmt.Println(superpower[0])
	fmt.Println(superpower[1:2])

	// slice和map复合字面值也可以使用&操作，但是go并没有提供自动解引用的功能

}
