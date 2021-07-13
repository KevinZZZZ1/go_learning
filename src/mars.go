package main

import "fmt"

func main() {
	fmt.Print("My weight on the surface of Mars is ")
	fmt.Print(120.0 * 0.3783)

	fmt.Print(" lbs, and I would be ")
	fmt.Print(27 * 365 / 678)
	fmt.Print(" years old")

	fmt.Println("My weight on the surface of Mars is", 120.0*0.3783, "lbs, and I would be", 27*365.2425/687, "years old.")

	fmt.Printf("My weight on the surface of Mars is %v lbs", 120.0 * 0.3783)
	fmt.Printf(" and I would be %v years old.\n", 27.0 * 365 / 687)

	fmt.Printf("My weight on the surface of %v is %v lbs.\n", "Earth", 120.0)

}
