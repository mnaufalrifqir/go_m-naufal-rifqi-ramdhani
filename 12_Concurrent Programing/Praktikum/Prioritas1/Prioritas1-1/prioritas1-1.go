package main

import (
	"fmt"
	"time"
)

func kelipatan(x int) {
	for i := 1; ; i++ {
		fmt.Println(i * x)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	var x int
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&x)
	go kelipatan(x)
	time.Sleep(time.Duration(x) * time.Second)
}
