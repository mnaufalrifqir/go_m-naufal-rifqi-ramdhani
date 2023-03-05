package main

import "fmt"

func playingDomino(cards [][]int, deck []int) interface{} {
	var req [][]int
	for _, value := range cards {
		if value[0] == deck[0] || value[0] == deck[1] || value[1] == deck[0] || value[1] == deck[1] {
			req = append(req, value)
		}
	}

	if len(req) != 0 {	//Tambahan untuk menentukan jumlah terbesar dari kedua sisi
		var max, idx int
		max = req[0][0] + req[0][1]
		idx = 0
		for key, value := range req {
			if value[0]+value[1] > max {
				idx = key
			}
		}
		return req[idx]
	}
	return "tutup kartu"
}

func main() {
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}}, []int{4, 3})) // [3, 4]
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 3}, {3, 4}, {2, 1}}, []int{3, 6})) // [6, 5]
	fmt.Println(playingDomino([][]int{{6, 6}, {2, 4}, {3, 6}}, []int{5, 1})) // "tutup kartu"
}
