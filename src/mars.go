package main

import "fmt"

func main() {

	// struct可以将多个不同类型的字段组合到一起

	// 这里是声明了一个location类型的struct，这样就能符用这个类型
	type location struct {
		lat  float64
		long float64
	}

	// 初始化struct，使用成对的key,value的形式初始化
	// 如果字段不设置初始化值的话，会使用字段类型的零值
	var spririt location = location{
		lat:  -1.9462,
		long: 354.4734,
	}

	// 初始化struct的另一种方式，没有显式指定字段名及其值，会以字段的声明顺序进行赋值
	var opportunity location = location{-17.5698, 175.45521}

	// 格式化打印struct
	fmt.Printf("%+v\n", opportunity)
	fmt.Printf("%v\n", spririt)


}
