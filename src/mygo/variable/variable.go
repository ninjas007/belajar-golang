package main

import (
	"fmt"
)

func main() {
	var firstName = "Tilis"
	var lastName = "Tiadi"
	const midName = "Adi"

	// pointer
	var numberA int = 4
	var numberB *int = &numberA
	*numberB = 5

	fmt.Println("Hello", firstName, midName, lastName)
	fmt.Println("Pointer", numberA, *numberB)
}
