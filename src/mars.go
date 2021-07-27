package main

import "fmt"

type celsius float64

type fahrenheit float64

type kelvin float64

func main() {

	var tmp celsius = 25

	fmt.Printf("%v celsius is %v in form of fahrenheit and also %v in form of kelvin", tmp, tmp.toFahrenheit(), tmp.toKelvin())

}

// func关键字 (接收者 接收者类型) 方法名(参数列表) 返回值类型
func (c celsius) toFahrenheit() fahrenheit {

	return fahrenheit((c*9)/5 + 32)

}

func (c celsius) toKelvin() kelvin {
	return kelvin(c + 272.15)
}
