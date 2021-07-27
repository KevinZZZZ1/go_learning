package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64

func main() {
	// go中的函数是头等的，可以用在整数，字符串或其他类型能用的地方
	// 		1. 函数可以赋值给变量
	//		2. 函数可以作为参数传递给其他函数
	//		3. 函数可以作为其他函数的返回值

	// 函数赋值给变量
	// var 变量名 func() 函数返回值
	var sensor func() kelvin = fakeSensor
	fmt.Println(sensor())

	sensor = realSensor
	fmt.Println(sensor())

	// 函数作为其他函数的参数
	measureTemperature(3, fakeSensor)

}

func measureTemperature(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%v°K \n", k)
		time.Sleep(time.Second)
	}
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0
}
