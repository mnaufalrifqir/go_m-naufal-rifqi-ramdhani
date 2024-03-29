package main

import "fmt"

func fibonacci(number int) int {
	if number == 0 || number == 1 {
		return number
	} else {
		return (fibonacci(number-1) + fibonacci(number-2))
	}
}

func main() {
	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144
}
