package main

import "fmt"

func main() {

	// 可以使用make函数对slice预分配容量，这样能避免切片底层对应数组扩容导致的内存分配和数组复制操作
	// 这表示创建一个slice，长度为0，容量为10
	var dwarfs []string = make([]string, 0, 10)
	// 如果是这种写法的话var dwarfs []string = make([]string, 10)
	// 表示切片的长度和容量都是10

	dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")

	dump("dwarfs", dwarfs)

	twoWorlds := terraform("new", "Venus", "Mars")
	fmt.Println(twoWorlds)

	var planets []string = []string{"Venus", "Mars", "Jupiter"}
	newPlanets := terraform("new", planets...)
	fmt.Println(newPlanets)

}

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)

}

// 函数的可变长度参数
func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))

	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}

	return newWorlds
}
