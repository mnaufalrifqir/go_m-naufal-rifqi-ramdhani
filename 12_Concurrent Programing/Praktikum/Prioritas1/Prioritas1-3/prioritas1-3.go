package main

import (
	"fmt"
)

func kelipatanTiga(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i * 3
	}
}

func main() {
	ch := make(chan int, 10)
	go kelipatanTiga(ch)
	for i := 1; i <= 10; i++ {
		fmt.Println(<-ch)
	}
}
