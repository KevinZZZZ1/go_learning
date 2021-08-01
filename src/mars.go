package main

import "fmt"

func main() {

	// struct可以将多个不同类型的字段组合到一起

	// 声明了一个变量curiosity，它的类型是struct
	// 这种声明类似Java中的匿名类
	var curiosity struct {
		lat  float64
		long float64
	}

	curiosity.lat = -4.5895
	curiosity.long = 137.4417

	fmt.Println(curiosity.lat, curiosity.long)
	fmt.Println(curiosity)

	// 这里是声明了一个location类型的struct，这样就能符用这个类型
	type location struct {
		lat  float64
		long float64
	}

	var spririt location
	spririt.lat = -14.5684
	spririt.long = 175.472636

	fmt.Println(spririt)

}
