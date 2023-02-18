package main

import "fmt"

func main() {
	var bilangan int

	fmt.Print("Masukkan bilangan yang akan dicek: ")
	fmt.Scanln(&bilangan)
	if bilangan%2 == 0 {
		fmt.Println(bilangan, "Merukapan bilangan genap")
	} else {
		fmt.Println(bilangan, "Merupakan bilangan ganjil")
	}
}
