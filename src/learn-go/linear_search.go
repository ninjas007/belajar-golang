package main

import "log"

func main() {
	arr := []int{6, 7, 5, 8, 9, 10, 1, 3, 2, 4}
	found, index := LinearSearch(9, arr...)
	log.Println(found, index)
}

func LinearSearch(value int, nums ...int) (found bool, index int) {

	for i, val := range nums {
		if val == value {
			found = true
			index = i
			return
		}
	}

	index = -1
	return
}