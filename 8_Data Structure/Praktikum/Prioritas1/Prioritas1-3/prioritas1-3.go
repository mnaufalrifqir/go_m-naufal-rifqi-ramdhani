package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	var hasil []int

	for i := 0; i < len(angka); i++ {
		var banyak int = 0
		for j := 0; j < len(angka); j++ {
			if angka[i] == angka[j] {
				banyak += 1
			}
		}
		if banyak == 1 {
			intVal, _ := strconv.Atoi(string(angka[i]))
			hasil = append(hasil, int(intVal))
		}
	}
	return hasil
}

func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6 3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}
