package main

import "log"

func main() {

	sorted := InsertionSort(6,5,3,1,8,7,2,4)
	log.Println("Sorted array", sorted)
}

func InsertionSort(nums ...int) (sorted []int) {
	for i := 0; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			log.Println(nums[j-1], ">", nums[j], "j=", j, "i=", i)

			// cara pertama
			// if nums[j-1] > nums[j] {
			// 	nums[j-1], nums[j] = nums[j], nums[j-1]
			// }

			// cara kedua
			if nums[j-1] > nums[j] {
				temp := nums[j]
				nums[j] = nums[j-1]
				nums[j-1] = temp	
			}
		}
	}

	return nums
}

