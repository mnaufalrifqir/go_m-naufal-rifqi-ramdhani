package main

import "fmt"

type Kelas []Student

type Student struct {
	Name  string
	Score float64
}

func (k Kelas) averageScore() float64 {
	var total float64
	for _, value := range k {
		total += float64(value.Score)
	}
	return total / 5
}

func (k Kelas) minScore() (string, float64) {
	var min float64 = k[0].Score
	var nama string = k[0].Name

	for _, value := range k {
		if value.Score < min {
			nama = value.Name
			min = value.Score
		}
	}
	return nama, min
}

func (k Kelas) maxScore() (string, float64) {
	var max float64 = k[0].Score
	var nama string = k[0].Name

	for _, value := range k {
		if value.Score > max {
			nama = value.Name
			max = value.Score
		}
	}
	return nama, max
}

func main() {
	var nama string
	var nilai float64
	var kelas Kelas

	for i := 1; i <= 5; i++ {
		fmt.Print("Input ", i, " Student's Name ")
		fmt.Scanln(&nama)
		fmt.Print("Input ", i, " Student's Score ")
		fmt.Scanln(&nilai)
		murid := Student{nama, nilai}
		kelas = append(kelas, murid)
	}
	fmt.Println("Average Score : ", kelas.averageScore())
	namaMin, scoreMin := kelas.minScore()
	fmt.Println("Min Score of Students :", namaMin, "(", scoreMin, ")")
	namaMax, scoreMax := kelas.maxScore()
	fmt.Println("Min Score of Students :", namaMax, "(", scoreMax, ")")
}
