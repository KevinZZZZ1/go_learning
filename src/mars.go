package main

import (
	"fmt"
	"math"
)

func main() {

	// 除了使用复合字面值来初始化map，还可以使用make函数来创建map
	// 这里的8表示预分配的空间
	var temperature map[float64]int = make(map[float64]int, 8)

	fmt.Println(len(temperature))

	var data []float64 = []float64{
		-28.0, -32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	var groups map[float64][]float64 = make(map[float64][]float64)

	for _, t := range data {
		g := math.Trunc(t/10) * 10
		groups[g] = append(groups[g], t)
	}

	// 这里的遍历不能严格保证顺序
	for key, value := range groups {
		fmt.Printf("%v: %v\n", key, value)
	}

}
