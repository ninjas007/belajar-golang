package main

import (
	"fmt"
)

func main() {
	// aritmatika
	var n1 int = 4
	var n2 int = 2
	var result = n1 * n2

	var nSiswa1 = 90
	var nSiswa2 = 50
	standarLulus := 60
	result2 := fmt.Sprintf("Nilai Standar: %d", standarLulus) + "\n"

	// perbandingan
	// operator logika (&&, ||, !)
	// operator perbandingan (<, >, ==, dll)
	if standarLulus < nSiswa1 {
		result2 += "Siswa1 Tidak Lulus \n"
	} else {
		result2 += "Siswa1 Lulus\n"
	}

	if standarLulus < nSiswa2 {
		result2 += "Siswa2 Tidak Lulus\n"
	} else {
		result2 += "Siswa2 Lulus\n"
	}

	fmt.Println(result)
	fmt.Println(result2)
}
