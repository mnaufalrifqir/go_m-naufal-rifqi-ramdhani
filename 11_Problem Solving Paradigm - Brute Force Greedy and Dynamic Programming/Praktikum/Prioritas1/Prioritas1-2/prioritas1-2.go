package main

import "fmt"

func segitigaPascal(numRows int) [][]int {
	var hasil = [][]int{}
	for i := 0; i < numRows; i++ {
		var temp = []int{}
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				temp = append(temp, 1)
			} else {
				temp = append(temp, hasil[i-1][j-1]+hasil[i-1][j])
			}
		}
		hasil = append(hasil, temp)
	}
	return hasil
}

func main() {
	fmt.Println(segitigaPascal(5))
}
