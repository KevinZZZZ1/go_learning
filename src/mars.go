package main

import (
	"fmt"
	"math/rand"
	"time"
)

var era = "AD"

func main() {
	// 有些地方不能使用var来声明变量，可以使用短声明来声明变量
	// 比如在for循环，if判断，switch中使用短声明

	// 这里有短声明来创建count变量
	for count := 0; count < 10; count++ {
		fmt.Println("count = ", count)
		time.Sleep(time.Second)
	}

	if count := rand.Intn(3); count == 0 {
		fmt.Println("Space Adventures")
	} else if count == 1 {
		fmt.Println("SpaceX")
	} else {
		fmt.Println("Virgin Galactic")
	}

	switch count := rand.Intn(5); count {

	case 0:
		fmt.Println("Space Adventure")

	case 1:
		fmt.Println("SpaceX")

	case 2:
		fmt.Println("Virgin Galactic")

	default:
		fmt.Println("Random spaceline #", count)

	}

}
