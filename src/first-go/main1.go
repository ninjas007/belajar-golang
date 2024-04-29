package main

import (
	"fmt"
	"net/http"
)

// untuk menampilkan ke website
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Index called ...")
	hello := "Hello World"
	w.Write([]byte(hello))
}

func coba(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Index called coba ...")
	hello := "Hello World coba"
	w.Write([]byte(hello))
}

// untuk melakukan routing
func handeRoute(){
    http.HandleFunc("/", index) // bisa begini
	// atau begini

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("Index called ...")
	// 	hello := "Hello World asdf"
	// 	w.Write([]byte(hello))
	// })

	http.HandleFunc("/coba", coba) 
	// http.HandleFunc("/coba", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("Index called ...")
	// 	hello := "Hello World coba"
	// 	w.Write([]byte(hello))
	// })
}

func main() {
    // port bebas, yang penting belum digunakan sama sekali.
    // saya rekomendasikan membuat port yang simple 
    // seperti 3000, 3002, 9000, 9090, 9999, dll..
    port := ":9999"
	startServer(port)
}

// untuk membuat web server
func startServer(port string) {
    handeRoute()
	fmt.Println("Server running at port ", port)
	http.ListenAndServe(port, nil)
}