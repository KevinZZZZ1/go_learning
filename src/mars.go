package main

import "fmt"

func main() {
	// 使用Printf对齐文本
	// 比如%4v: 表示向左填充4个宽度
	fmt.Printf("%-15v $%4v\n", "SpaceX", 84)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)

}
