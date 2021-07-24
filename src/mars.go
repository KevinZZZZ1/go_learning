package main

import "fmt"

func main() {
	// 常量(const)和变量(var)

	const lightSpeed = 299792 // km/s
	var distance = 5000000    // km

	fmt.Println(distance/lightSpeed, "seconds")

	distance = 4000000
	fmt.Println(distance/lightSpeed, "seconds")

}
