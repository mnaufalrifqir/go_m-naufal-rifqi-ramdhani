package main

import "fmt"

func caesar(offset int, input string) string {
	var hasil string

	for i := 0; i < len(input); i++ {
		hasil += string(rune((int(input[i])-97+offset)%26 + 97))
	}

	return hasil
}

func main() {
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // def
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
