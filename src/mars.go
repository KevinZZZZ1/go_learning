package main

import (
	"fmt"
)

type person struct {
	name, superpower string
	age              int
}

// 指针作为函数参数
func birthday(p *person) {
	p.age++
}

// 指针作为方法接收者
func (p *person) birthday() {
	p.age++
}

type stats struct {
	level             int
	endurance, health int
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

type character struct {
	name  string
	stats stats
}

func main() {
	// go中函数和方法都是按值传递的，也就是说它们操作的是被传递参数的副本
	var terry *person = &person{
		name: "Terry",
		age:  15,
	}

	terry.birthday()
	fmt.Printf("%+v\n", terry)

	birthday(terry)
	fmt.Printf("%+v\n", terry)

	player := character{name: "Matthias"}
	// 获得结构体中指定字段的地址
	levelUp(&player.stats)

	fmt.Printf("%+v\n", player.stats)
}
