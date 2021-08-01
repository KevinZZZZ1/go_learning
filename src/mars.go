package main

import (
	"fmt"
)

func main() {

	var administrator *string

	scolese := "Christopher J. Scolese"
	administrator = &scolese
	fmt.Println(*administrator)
	// 输出Christopher J. Scolese

	bolden := "Charles F. Bolden"
	administrator = &bolden
	fmt.Println(*administrator)
	// 输出Charles F. Bolden

	bolden = "Charles Frank Bolden Jr."
	fmt.Println(*administrator)
	// 输出Charles Frank Bolden Jr.

	*administrator = "Maj. Gen. Charles Frank Bolden Jr."
	fmt.Println(bolden)
	// 输出Maj. Gen. Charles Frank Bolden Jr.

	major := administrator
	*major = "Major General Charles Frank Bolden Jr."
	fmt.Println(bolden)
	// 输出Major General Charles Frank Bolden Jr.

	// 如果两个指针类型的变量中存储的地址是一样的，它们的比较就是true
	fmt.Println(administrator == major)
	// 输出true

	lightfoot := "Robert M. Lightfoot Jr."
	administrator = &lightfoot
	fmt.Println(administrator == major)
	// 输出false

	charles := *major
	*major = "Charles Bolden"
	fmt.Println(charles)
	// 输出Major General Charles Frank Bolden Jr.
	fmt.Println(bolden)
	// 输出Charles Bolden

	charles = "Charles Bolden"
	fmt.Println(charles == bolden)
	// 输出true
	fmt.Println(&charles == &bolden)
	// 输出false
}
