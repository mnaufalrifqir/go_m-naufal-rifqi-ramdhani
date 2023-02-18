package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var kata string
	var valid bool = true

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan kata: ")
	scanner.Scan()
	kata = scanner.Text()
	fmt.Println("Captured: ", kata)
	for i := 0; i < (len(kata) / 2); i++ {
		if kata[i] != kata[len(kata)-(i+1)] {
			fmt.Println("Tidak Palindrome")
			valid = false
			break
		}
	}
	if valid == true {
		fmt.Println("Palindrome")
	}
}
