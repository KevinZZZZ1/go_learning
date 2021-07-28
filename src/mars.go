package main

import "fmt"

// 数组是一种固定长度的有序元素集合

func main() {
	// 数组的声明，这里要指定数组的长度
	var planets [8]string

	// 数组元素的赋值
	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"

	// 数组元素的访问
	earth := planets[2]
	fmt.Println(earth)

	// 使用len()内置函数得到数组的长度
	fmt.Println(len(planets))
	// 对于数组中未赋值的元素，会使用数组类型的零值来作为默认值
	fmt.Println(planets[3] == "")

	// 使用复合字面值声明和初始化数组
	var dwarfs [5]string = [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Println(dwarfs)
	// 在复合字面值中使用...作为数组的长度，go编译器会计算出数组的元素数量
	var dwarfs2 = [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Println(dwarfs2)

	// 使用for循环遍历数组
	for i := 0; i < len(dwarfs); i++ {
		fmt.Println(dwarfs[i])
	}

	// 使用range关键字遍历数组
	for i, dwarf := range dwarfs {
		fmt.Println(i, dwarf)
	}

}
