package main

import "fmt"

var hasilFibonacci = map[int]int{}

func fibo(n int) int {
	value, exist := hasilFibonacci[n]
	if exist {
		return value
	} else if n == 0 || n == 1 {
		hasilFibonacci[n] = n
		return n
	} else {
		hasilFibonacci[n] = fibo(n-2) + fibo(n-1)
		return hasilFibonacci[n]
	}
}

func main() {
	fmt.Println(fibo(0))  //0
	fmt.Println(fibo(1))  //1
	fmt.Println(fibo(2))  //1
	fmt.Println(fibo(3))  //2
	fmt.Println(fibo(5))  //5
	fmt.Println(fibo(6))  //8
	fmt.Println(fibo(7))  //13
	fmt.Println(fibo(9))  //34
	fmt.Println(fibo(10)) //55
}
