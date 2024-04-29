package main

import "fmt"

func main() {
	chicken := map[string]int{}

	chicken["januari"] = 10
	chicken["februari"] = 20

	fmt.Println("Januari", chicken["januari"])
	fmt.Println("Februari", chicken["asfd"])
}
