package main

import "fmt"

func main() {
	var nilai int8

	fmt.Print("Masukkan nilai: ")
	fmt.Scanln(&nilai)

	if nilai >= 80 && nilai <= 100 {
		fmt.Println("A")
	} else if nilai >= 65 && nilai < 80 {
		fmt.Println("B")
	} else if nilai >= 50 && nilai < 65 {
		fmt.Println("C")
	} else if nilai >= 35 && nilai < 50 {
		fmt.Println("D")
	} else if nilai >= 0 && nilai < 35 {
		fmt.Println("E")
	} else {
		fmt.Println("Nilai Invalid")
	}
}
