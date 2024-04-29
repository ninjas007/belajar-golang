package main

import (
	"fmt"
	"strings"
)

func main() {
	var names = []string{"Aco", "Adi"}
	luas, keliling := calculate(2)
	total := hitung(2, 4, 6, 8, 10)

	getMessage("Hello", names)

	fmt.Println("Rumus Persegi")
	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println("Total", total)

	closure()

	callback(1)
}

func getMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")

	fmt.Println(message, nameString)
}

// multiple return
func calculate(num int) (int, int) {
	luas := num * num   // luas persegi
	keliling := 4 * num // keliling persegi

	return luas, keliling
}

// variadic
func hitung(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += int(number)
	}

	// for _, number := range numbers {
	// 	total += number
	// }

	// avg := int(total) / int(len(numbers))

	return total
}

func closure() {
	getMinMax := func(n []int) (int, int) {
		var min, max int

		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}

		return min, max
	}

	numbers := []int{5, 3, 1, 5, 6}
	min, max := getMinMax(numbers)

	fmt.Println("nilai min", min)
	fmt.Println("nilai max", max)
}

func callback(i int) {
	hasil := filterData([]string{"anton", "andi", "budi"}, func(each string) bool {
		return true
	})

	fmt.Println("hasil", hasil)
	fmt.Println("hasil", hasil[i])
}

func filterData(data []string, callback func(string) bool) []string {
	var result []string

	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}

	return result
}
