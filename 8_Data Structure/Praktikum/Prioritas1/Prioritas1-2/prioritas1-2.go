package main

import "fmt"

func Mapping(slice []string) map[string]int {
	var hasil = map[string]int{}

	//karena tipe data primitive(int) cukup seperti ini
	for i := 0; i < len(slice); i++ {
		hasil[slice[i]]++
	}
	return hasil
}

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"})) //map[adi:1 asd:2 qwe:3]
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))                      //map[asd:2 qwe:1]
	fmt.Println(Mapping([]string{}))                                         //map[]
}
