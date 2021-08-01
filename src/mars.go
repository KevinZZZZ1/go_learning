package main

import "fmt"

type location struct {
	lat  float64
	long float64
}

func main() {

	var spririt location = location{
		-4.5895,
		-137.4417,
	}

	// 这里将struct类型进行赋值，会创建一个新的对象赋值给bradbury，所以在修改spririt的时候bradbury并不会被修改
	bradbury := spririt

	spririt.long += 1

	fmt.Println(spririt, bradbury)

	curiosity := bradbury

	addLong(bradbury)

	fmt.Println(bradbury, curiosity)

}

// 这里传入一个location类型的struct，这里的参数l也是函数实际传入参数的一个复制，即和传入参数是相互独立的
func addLong(l location) {
	l.long += 10
}
