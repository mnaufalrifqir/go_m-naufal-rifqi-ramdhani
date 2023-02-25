package main

import "fmt"

func pow(x int, n int) int {
	hasil := 1
	for i := n; i > 0; i /= 2 { //O(log n)
		if i%2 == 1 {
			hasil *= x
		}
		x *= x
	}
	return hasil
}

func main() {
	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(5, 3))  // 125
	fmt.Println(pow(10, 2)) // 100
	fmt.Println(pow(2, 5))  // 32
	fmt.Println(pow(7, 3))  // 343
}
