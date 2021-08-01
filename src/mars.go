package main

import (
	"fmt"
	"strings"
)

// 定义一个接口变量
var t interface {
	talk() string
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

// 定义了一个接口类型
type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func main() {

	// 接口主要是定义了可以做什么，接口可以通过列举类型必须满足的一组方法来进行声明
	// GO中不需要显式声明实现了什么接口

	t = martian{}
	fmt.Println(t.talk())

	t = laser(3)
	fmt.Println(t.talk())

	shout(martian{})
	shout(laser(5))

}
