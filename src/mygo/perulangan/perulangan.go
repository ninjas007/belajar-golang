package main

import "fmt"

func main() {
	count := 10
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			if i > j {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}
