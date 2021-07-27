package main

import "fmt"

func main() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)

	fmt.Print("kelvin is ", kelvin, ", celsius is ", celsius)
}

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}
