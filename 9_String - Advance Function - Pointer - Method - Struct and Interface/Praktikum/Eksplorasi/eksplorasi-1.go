package main

import "fmt"

type student struct {
	name       string
	nameEncode string
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	const offset int = 100
	var nameEncode = ""

	for i := 0; i < len(s.name); i++ {
		nameEncode += string(rune((int(s.name[i])-97+offset)%26 + 97))
	}
	s.nameEncode = nameEncode
	return nameEncode
}

func (s *student) Decode() string {
	const offset int = 100
	var nameDecode = ""

	for i := 0; i < len(s.nameEncode); i++ {
		huruf := (int(s.nameEncode[i]) - (offset % 26))
		for huruf < int('a') {
			huruf += 26
		}
		nameDecode += string(rune(huruf))
	}
	s.name = nameDecode
	return nameDecode
}

func main() {
	var menu int
	var a student = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu? ")
	fmt.Scan(&menu)

	if menu == 1 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.name)
		fmt.Print("\nEncode of student's name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student Name: ")
		fmt.Scan(&a.nameEncode)
		fmt.Print("\nDecode of student's name " + a.nameEncode + " is : " + c.Decode())
	}
}
