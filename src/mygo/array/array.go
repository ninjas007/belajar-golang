package main

import "fmt"

func main() {
	fruits := [4]string{"apple", "manggo", "orange", "melon"}
	length := len(fruits)

	fmt.Println("Semua buah", fruits)
	for i := 0; i < length; i++ {
		fmt.Println("Buah", fruits[i])
	}
}
