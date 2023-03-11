package main

import (
	"fmt"
	"sync"
)

func main() {
	var frekuensi = map[rune]int{}
	var frekuensi2 = map[rune]int{}
	var wg sync.WaitGroup

	kalimat := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"
	wg.Add(2)
	go func(i, j int) {
		defer wg.Done()
		for _, value := range kalimat[i:j] {
			if value >= 65 && value <= 90 {
				frekuensi[rune(int(value)+32)]++
			} else if value >= 97 && value <= 122 {
				frekuensi[value]++
			}
		}
	}(0, len(kalimat)/2)

	go func(i, j int) {
		defer wg.Done()
		for _, value := range kalimat[i:j] {
			if value >= 65 && value <= 90 {
				frekuensi2[rune(int(value)+32)]++
			} else if value >= 97 && value <= 122 {
				frekuensi2[value]++
			}
		}
	}(len(kalimat)/2, len(kalimat))

	wg.Wait()
	for key, value := range frekuensi2 {
		frekuensi[key] += value
	}
	for key, value := range frekuensi {
		fmt.Printf("%s : %d\n", string(key), value)
	}
}
