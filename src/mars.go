package main

import "fmt"

func reclassify(planets *[]string) {
	*planets = (*planets)[0:8]
}

func main() {
	// map在赋值或被作为参数传递给函数（方法）的时候不会被复制，而是以指针的形式传递

	// slice在指向数组元素的时候也使用了指针
	// 每个slice内部都会被表示为一个包括3个元素的结构：
	// 	数组的指针
	//	slice容量
	// 	slice长度
	// slice被直接传递给方法（函数）时，slice内部的指针就可以对底层的数组进行操作
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
		"Pluto",
	}

	reclassify(&planets)
	fmt.Println(planets)
}
