package main

import "fmt"

type Pair struct {
	Name  string
	Count int
}

func MostAppearItem(items []string) []Pair {
	var hasil = map[string]int{}
	var arrPair []Pair

	for i := 0; i < len(items); i++ {
		hasil[items[i]]++
	}

	for key, value := range hasil {
		kata := Pair{key, value}
		arrPair = append(arrPair, kata)
	}

	for i := 0; i < len(arrPair)-1; i++ {
		for j := i + 1; j < len(arrPair); j++ {
			if arrPair[i].Count > arrPair[j].Count {
				arrPair[i], arrPair[j] = arrPair[j], arrPair[i]
			}
		}
	}

	return arrPair
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
}
