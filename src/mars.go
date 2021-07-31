package main

import "fmt"

func main() {

	var planets map[string]string = map[string]string{
		"Earth": "Sector ZZ9",
		"Mars":  "Sector ZZ9",
	}

	// map在赋值的时候并不会新建副本赋值，而是直接将对象赋值过去
	planetsMarkII := planets

	// 这里会把planets和planetsMarkII的Key为Earth的Value都设置为whoops
	planets["Earth"] = "whoops"
	fmt.Println(planets)
	fmt.Println(planetsMarkII)

	// 在map中删除对应key
	delete(planets, "Earth")
	fmt.Println(planetsMarkII)
	fmt.Println(planets)

}
