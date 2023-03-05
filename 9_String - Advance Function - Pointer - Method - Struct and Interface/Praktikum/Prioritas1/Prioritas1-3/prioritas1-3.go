package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	var hasil string
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a)+1; j++ {
			temp := a[i:j]
			if strings.Contains(b, temp) && len(temp) > len(hasil) {
				hasil = temp
			}
		}
	}
	return hasil
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
}
