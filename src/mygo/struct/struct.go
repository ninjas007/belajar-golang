package main

import "fmt"

type Student struct {
	name  string
	grade int
}

func main() {
	var mhs Student

	mhs.name = "Adi"
	mhs.grade = 90

	fmt.Println("name", mhs.name)
	fmt.Println("grade", mhs.grade)
}
