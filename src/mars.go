package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type location struct {
	Name string
	Lat  float64
	Long float64
}

func main() {

	// 声明一个struct slice
	var locations []location = []location{
		{Name: "Bradbury Landing", Lat: -4.5895, Long: 175.4513541},
		{Name: "Columbia Memorial Station", Lat: -14.4551, Long: 137.41654},
		{Name: "Challenger Memorial Statition", Lat: -1.34265, Long: 354.454},
	}

	fmt.Println(locations)

	// struct格式化为json
	var curiosity location = location{
		Lat:  -4.2151,
		Long: 137.4121,
		Name: "ceeeeb",
	}

	// 使用Marshal函数时要求struct中的字段都是大写开头的，即可以导出的
	// 如果是不可以导出的，将会输出{}
	bytes, err := json.Marshal(curiosity)

	exitOnError(err)

	fmt.Println(string(bytes))

	// 如果想自定义json输出的字段key的话，可以使用struct标签来自定义
	type people struct {
		Name    string `json:"user_name"`
		Age     int    `json:"user_age"`
		Address string `json:"user_addr"`
	}

	var user people = people{
		Name:    "kvein",
		Age:     30,
		Address: "USA",
	}

	bytes, err = json.Marshal(user)

	exitOnError(err)

	fmt.Println(string(bytes))

}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
