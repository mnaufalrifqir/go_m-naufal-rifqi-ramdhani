package main

import ( //Tanpa packages
	"fmt"
)

func decToBiner(angka int) int { 
	var temp []int 
	for angka > 0 {
		if angka%2 == 0 {
			temp = append(temp, 0)
		} else {
			temp = append(temp, 1)
		}
		angka /= 2
	}
	result := 0
	for i := len(temp) - 1; i >= 0; i-- {
		pangkat := 1
		for tempI := i; tempI > 0; tempI-- {
			pangkat *= 10
		}
		result += temp[i] * pangkat
	}
	return result
}
func arrBiner(angka int) []int {
	var hasil []int

	hasil = append(hasil, 0)
	for i := 1; i <= angka; i++ {
		hasil = append(hasil, decToBiner(i))
	}
	return hasil
}

func main() {
	fmt.Println(arrBiner(2))
	fmt.Println(arrBiner(3))
	fmt.Println(arrBiner(4))
}
