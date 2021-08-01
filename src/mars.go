package main

import "fmt"

func main() {

	// 指针是指向另一个变量地址的变量

	// &表示地址操作符，通过&可以获得变量的内存地址
	// *和&的作用相反，是用来解引用的，表示得到内存地址指向的值

	var answer int32 = 42
	fmt.Println(&answer)

	// 定义一个指向int32类型的指针，*放在变量前面表示解引用操作，放在类型前面表示指针类型
	var address *int32 = &answer
	fmt.Printf("%T\n", address)
	fmt.Println(*address)

}
