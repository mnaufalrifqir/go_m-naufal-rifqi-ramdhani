package main

import (
	"fmt"
)

func intToRomawi(number int) string {
	var romawi = [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var decimal = [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	var hasil string = ""

	for i := 0; i < len(romawi) && number > 0; i++ {
		for number >= decimal[i] {
			number -= decimal[i]
			hasil += romawi[i]
		}
	}

	return hasil
}

func main() {
	fmt.Println(intToRomawi(4))
	fmt.Println(intToRomawi(9))
	fmt.Println(intToRomawi(23))
	fmt.Println(intToRomawi(2021))
	fmt.Println(intToRomawi(1646))
}
