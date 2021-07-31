package main

import "fmt"

func main() {

	// 声明一个<string, int>的map
	var temperature map[string]int = map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}

	// 获得Key为Earth对应的Value
	temp := temperature["Earth"]
	fmt.Printf("On average the Earth is %v.\n", temp)

	// 设置Key为Earth，Value为16到map中
	temperature["Earth"] = 16
	temperature["Venus"] = 464
	fmt.Println(temperature)

	moon := temperature["Moon"]
	fmt.Println(moon)
}
