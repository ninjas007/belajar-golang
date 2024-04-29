package main

import "log"

func main() {
	num := 10
	for i := 0; i < num; i++ {
		log.Println(fibo(i))
	}
}

func fibo(n int) int {
	// index ke 0 dan index ke 1 hasilnya adalah
	if n <= 1 {
		return 1
	}

	// masuk index kedua maka 2 - 1 + 2 - 2
	return fibo(n-1) + fibo(n-2)
}