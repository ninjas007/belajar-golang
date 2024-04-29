package main

import "fmt"

func main() {
	contohPointer()
}

func contohPointer() {
	nama := "NooBee"
    namaPointer := &nama

    fmt.Println(nama) // NooBee
    fmt.Println(namaPointer) 
}