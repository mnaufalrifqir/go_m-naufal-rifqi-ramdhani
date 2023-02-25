package main

import "fmt"

func hitungSelisihDiagonal(matriks [][]int) int {
	var d1, d2 int
	for i := 0; i < len(matriks); i++ {
		for j := 0; j < len(matriks[i]); j++ {
			if i == j {
				d1 += matriks[i][j]
			}
			if j == len(matriks[i])-(i+1) {
				d2 += matriks[i][len(matriks[i])-(i+1)]
			}
		}
	}
	if d1 > d2 {
		return d1 - d2
	} else {
		return d2 - d1
	}
}

func main() {
	var matriks [][]int = [][]int{{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9}}
	fmt.Println(hitungSelisihDiagonal(matriks))
}
