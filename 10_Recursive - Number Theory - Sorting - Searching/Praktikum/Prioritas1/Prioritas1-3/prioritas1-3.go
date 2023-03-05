package main

import (
	"fmt"
	"math"
)

func primeX(number int) int {
	num := 1
	urutan := 0

	for urutan < number {
		var check bool = true
		if num <= 1 {
			check = false
		}
		for i := 2; i <= int(math.Sqrt(float64(num))); i++ { //O(sqrt(n))
			if num%i == 0 {
				check = false
			}
		}

		if check {
			urutan++
		}
		num++
	}
	return num - 1
}

func main() {
	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
	fmt.Println(primeX(10))
}
