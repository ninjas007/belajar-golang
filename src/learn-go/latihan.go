// Silahkan cari :

// Median (nilai tengah, harus di urutkan)
// Max => Nilai paling tinggi. Kalau ada lebih dari 1, silahkan tampilkan semuanya
// Min => Nilai paling kecil. Kalau ada lebih dari 1, silahkan tampilkan semuanya
// Average => rata rata
// Cari nilai yang diatas 80 (Binary Search)

package main

import (
	"fmt"
	"sort"
	"log"
)

type Examination struct{
	Name string
	Score int
}

func main() {
	var examResult = []Examination{
		{Name: "Budi", Score: 90},
		{Name: "Andi",Score: 85},
		{Name: "Nayla",Score: 87},
		{Name: "Danu",Score: 80},
		{Name: "Rahmat",Score: 75},
		{Name: "Erika",Score: 83},
		{Name: "Siska",Score: 87},
		{Name: "Mita",Score: 94},
		{Name: "Shinta",Score: 82},
		{Name: "Jojo",Score: 83},
	}

	var length = len(examResult)

	// sorting by score
	sortByScore(examResult) // [{Rahmat 75} {Danu 80} {Shinta 82} {Erika 83} {Jojo 83} {Andi 85} {Nayla 87} {Siska 87} {Budi 90} {Mita 94}]

	fmt.Println("Median: ", med(examResult, length)) // expected 85
	fmt.Println("Maksimal: ", max(examResult, length)) // expected 94
	fmt.Println("Minimal: ", min(examResult, length)) // expected 75
	fmt.Println("Average: ", avg(examResult, length)) // expected 84
	fmt.Println("Score diatas 80: ", filterByScore(examResult, length, 80))
}

func sortByScore(examResult []Examination) {
	// sort by score
	// for i := 0; i < length; i++ {
	// 	for j := i; j < length; j++ {
	// 		if examResult[i].Score > examResult[j].Score {
	// 			examResult[i], examResult[j] = examResult[j], examResult[i]
	// 		}
	// 	}
	// }
	
	// or using this to sorting by score
	sort.Slice(examResult, func(i, j int) bool {
		return examResult[i].Score < examResult[j].Score
	})

	log.Println(examResult)
}

func med(examResult []Examination, length int) int {
	return examResult[length/2].Score
}

func max(examResult []Examination, length int) int {
	return examResult[length-1].Score
}

func min(examResult []Examination, length int) int {
	return examResult[0].Score
}

func avg(examResult []Examination, length int) int {
	resultScore := 0;
	for i := 0; i < length; i++ {
		resultScore += examResult[i].Score
	}

	return resultScore / length;
}

func filterByScore(examResult []Examination, length int, target int) int {
	
}