package main

import "fmt"

type coordinate struct {
	d, m, s float64
	h       rune
}

// go中没有Java中的构造方法，但是可以用一个函数用来构造struct，这样就能让我们在创建struct的时候同步做些其他的事情
func newCoordinate() coordinate {
	return coordinate{}
}

func main() {

	// go中没有class，对象，继承
	// 我们可以通过struct和方法实现面向对象，即我们可以使用方法关联到类型的方式实现

	lat := coordinate{4, 35, 22.2, 'S'}

	long := coordinate{137, 26, 30.12, 'E'}

	coor := newCoordinate()

	fmt.Println(lat.decimal(), long.decimal(), coor.decimal())

}

func (c coordinate) decimal() float64 {
	var sign float64 = 1
	switch c.h {
	case 'S', 's', 'E', 'e':
		sign = -1
	}

	return sign * (c.d + c.m/60 + c.s/3600)
}
