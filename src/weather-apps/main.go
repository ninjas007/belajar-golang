package main

import (
	"log"
	"net/http"
	"weather-apps/app"
)

func main() {
	/*
	   bebas mau dibuat berapa port nya
	   example : 4444,4000,8080,8000, ...etc
	   yang penting harus belum pernah digunakan
	*/
	port := ":4444"

	// ======= UPDATE =======
	// daftarkan handler tadi ke endpoint /weather
	http.HandleFunc("/weather", app.GetCurrentWeather)
	// ======================

	// get users
	http.HandleFunc("/users", app.GetUsers)

	// tampilkan log sebagai penanda bahwa server
	// telah running
	log.Println("server running at port", port)

	// proses running server
	http.ListenAndServe(port, nil)
}
