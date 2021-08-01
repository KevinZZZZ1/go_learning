package main

import "fmt"

type person struct {
	age int
}

func (p *person) birthday() {
	if p == nil {
		return
	}
	p.age++
}

func main() {
	// go中nil表示“无”或“零”，是一个零值
	// 指针，slice，map，函数类型变量和接口的零值都是nil
	// 尝试解引用一个nil的指针会导致panic
	var nobody *person
	fmt.Println(nobody)

	// 如果p是nil的话，在给p解引用的时候才会panic，所以单纯调用nobody.birthday()并不是panic
	nobody.birthday()

	var fn func(a, b int) int
	fmt.Println(fn == nil)

	// range len() append()可以正确处理nil的slice和map
	var soup []string

	for _, ingredient := range soup {
		fmt.Println(ingredient)
	}

	fmt.Println(len(soup))

	soup = append(soup, "onion", "carrot")
	fmt.Println(soup)

	var soups map[string]int
	fmt.Println(soup == nil)

	mesurement, ok := soups["onion"]
	if ok {
		fmt.Println(mesurement)
	}

	for ingredient, mesurement := range soups {
		fmt.Println(ingredient, mesurement)
	}

	// 未被赋值的接口变量，其接口类型和值都是nil，变量本身也是nil
	// 只有当类型和值都是nil的时候，接口变量才等于nil
	var test interface{}
	// 分别打印出类型，值，是否等于nil
	fmt.Printf("%T, %v %v\n", test, test, test == nil)
	var p *int
	test = p
	fmt.Printf("%T, %v %v\n", test, test, test == nil)
	// 输出接口变量的内部表示
	fmt.Printf("%#v\n", test)

}
