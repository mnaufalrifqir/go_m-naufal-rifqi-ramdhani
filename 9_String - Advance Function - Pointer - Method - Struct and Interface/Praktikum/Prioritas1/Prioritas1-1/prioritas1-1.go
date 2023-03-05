package main

import "fmt"

type Car struct {
	Type   string
	FuelIn float64
}

func (c Car) hitungJarak() float64 {
	return c.FuelIn / 1.5 // karena pada aturannya 1.5 liter dapat menempuh 1 mil atau 1.5L/mil
}

func main() {
	car1 := Car{"Jeep", 30.0}
	car2 := Car{"Kijang", 18.0}
	car3 := Car{"Sedan", 12.0}

	fmt.Println(car1.hitungJarak(), "mil")
	fmt.Println(car2.hitungJarak(), "mil")
	fmt.Println(car3.hitungJarak(), "mil")
}
