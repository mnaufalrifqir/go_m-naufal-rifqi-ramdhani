package main

import "fmt"

func main() {
	var a, b, t, luas float64

	fmt.Print("Masukkan panjang sisi atas: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan panjang sisi bawah: ")
	fmt.Scanln(&b)
	fmt.Print("Masukkan tinggi: ")
	fmt.Scanln(&t)
	luas = 0.5 * (a + b) * t
	fmt.Println("Luas Trapesium: ", luas)
}
