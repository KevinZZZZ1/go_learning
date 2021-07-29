package main

import (
	"fmt"
	"strings"
)

type StringSlice []string

func (s StringSlice) Print() {
	for i := range s {
		s[i] = s[i] + " new"
	}
}

func main() {

	// slice（切片）是指向数组的窗口，即是数组的一个视图，不会导致数组被修改
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	// 这里就是创建了一个切片，表示的区间是[0, 4)
	var terrestrial []string = planets[0:4]
	var gasGiants []string = planets[4:6]
	var iceGiants []string = planets[6:8]

	fmt.Println(terrestrial, gasGiants, iceGiants)

	// 同时也可以进行字符串的切分
	var neptune string = "neptune"
	var tune string = neptune[3:]

	fmt.Println(tune)

	neptune = "Poseidon"
	fmt.Println(tune)

	hyperspace(planets[:])

	fmt.Println(strings.Join(planets[:], ", "))

	var ss StringSlice = planets[2:4]
	ss.Print()
	fmt.Println(ss)

}

// 切片的长度不是类型的一部分，所以这里可以是任意长度的切片
func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}
