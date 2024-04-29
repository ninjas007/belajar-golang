package main

import "fmt"

func main() {
	fruits := make([]string, 2)
	fruits[0] = "Apple"
	fruits[1] = "Manggo"

	fmt.Println("Buah", fruits[1])
}
