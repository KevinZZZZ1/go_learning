package main

import "fmt"

func main() {

	fmt.Println("There is a cavern entrance here and a path to the east.")

	var command = "go inside"

	// 这里的变量既可以是字符串也可以是数字
	switch command {

	case "go east":
		fmt.Println("You head further up the mountain.")

	case "enter cave", "go inside":
		fmt.Println("You find yourself in a dimly lit cavern.")

	case "read sign":
		fmt.Println("The sign read 'No Minors'.")

	default:
		fmt.Println("Didn't quite get that.")

	}

	var room = "lake"

	switch {

	case room == "cave":
		fmt.Println("You find yourself in a dimly lit cavern.")

	case room == "lake":
		fmt.Println("The ice seems solid enough.")
		// go 中的case语句是不用break的，这一点和Java中的不一样，
		// failthrough 关键字，表达的意思是继续执行下一个case，类似Java中没有break的switch语句
		fallthrough

	case room == "underwater":
		fmt.Println("The water is freezing cold.")

	case room == "default":
		fmt.Println("default value.")
	}

}
