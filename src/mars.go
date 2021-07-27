package main

import "fmt"

// 因为函数字面值（即匿名函数）要保留外部作用域的变量引用，所以函数字面值都是闭包的
// 闭包就是由于匿名函数封闭并包围作用域中的变量而得名的

type kelvin float64

type sensor func() kelvin

func main() {
	var sensor sensor = calibrate(realSensor, 5)
	fmt.Println(sensor())
}

func realSensor() kelvin {
	return 0
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		// 这里能获取到当前匿名函数作用域之外的参数，即保留了外部作用域的变量
		return offset + s()
	}
}
