package main

import "fmt"

func main() {

	var planets []string = []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	// 这里切片中的第三个参数表示切片的容量，会对应新建一个长度为4的地层数组，如果使用append等函数时容量不够的话，就会进行扩容（扩大为原来容量的2倍）
	terrestrial := planets[0:4:4]
	worlds := append(terrestrial, "Ceres")

	// 这里对应三个不同的地层数组
	dump("plants", planets)
	dump("terrestrial", terrestrial)
	dump("worlds", worlds)

	terrestrial = planets[0:4]
	// 这里会把plants切片中第5个元素替换为Ceres
	worlds = append(terrestrial, "Ceres")

	dump("plants", planets)
	dump("terrestrial", terrestrial)
	dump("worlds", worlds)
}

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}
