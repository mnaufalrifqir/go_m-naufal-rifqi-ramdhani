package main

import (
	"fmt"
	"sync"
	"time"
)

type Number struct {
	angka int
	m     sync.Mutex
}

func (n *Number) set(x int) {
	n.m.Lock()
	defer n.m.Unlock()
	n.angka = x
}

func (n *Number) get() int {
	n.m.Lock()
	defer n.m.Unlock()
	return n.angka
}

func faktorial(n int) int {
	var result = Number{}
	go func() {
		result.set(1)
		for i := n; i > 0; i-- {
			result.set(result.get() * i)
		}
	}()
	time.Sleep(time.Second)
	return result.get()
}

func main() {
	fmt.Println(faktorial(5))
}

//RACE CONDITION
// func faktorial(n int) int {
// 	result := 1
// 	go func() {
// 		for i := n; i > 0; i-- {
// 			result *= i
// 		}
// 	}()
// 	return result
// }

// func main() {
// 	fmt.Println(faktorial(5))
// }
