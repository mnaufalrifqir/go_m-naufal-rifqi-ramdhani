package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	var newArray []string
	var checkDup bool

	for i := 0; i < len(arrayA); i++ {
		for j := 0; j < len(newArray); j++ {
			if arrayA[i] == newArray[j] {
				checkDup = true
				break
			}
		}
		if !checkDup {
			newArray = append(newArray, arrayA[i])
		}
		checkDup = false
	}

	for k := 0; k < len(arrayB); k++ {
		for l := 0; l < len(newArray); l++ {
			if arrayB[k] == newArray[l] {
				checkDup = true
				break
			}
		}
		if !checkDup {
			newArray = append(newArray, arrayB[k])
		}
		checkDup = false
	}
	return newArray
}

func main() {
	// Test cases

	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	// ["king", "devil jin", "akuma", eddie", "steve", "geese"]

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	// ["sergei", "jin", "steve", "bryan"]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	// ["alisa", "yoshimitsu", "devil jin", "law"]

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	// ["devil jin", "sergei"]

	fmt.Println(ArrayMerge([]string{}, []string{}))

	// []
}
