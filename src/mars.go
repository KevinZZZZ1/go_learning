package main

import (
	"fmt"
	"sort"
)

func main() {

	// go中没有set数据类型，可以使用map来创建一个类似set的能力
	var temperatures []float64 = []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	set := make(map[float64]bool)

	for _, t := range temperatures {
		set[t] = true
	}

	// 已递增的顺序输出set中的value

	// 这里创建一个slice用于存放集合中元素，且方便排序
	unique := make([]float64, 0, len(set))

	for t := range set {
		if set[t] {
			unique = append(unique, t)
		}
	}

	sort.Float64s(unique)

	fmt.Println(unique)

}
