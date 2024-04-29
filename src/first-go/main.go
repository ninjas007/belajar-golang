package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string
	Class string
}

func main() {
	var users []User
	fmt.Println("User saat ini : ", len(users))

	// lalu ada data yang dikirim dari Frontend
	dataDariFrontend := `{"Name": "NooBee", "Class": "B"}`

	// mengubah string menjadi []byte
	data := []byte(dataDariFrontend)

	// menyiapkan object yang akan dimasuki data dari frontend
	var user User

	// proses parsing dari []byte ke object struct. dilakukan menggunakan pointer (&)
	// proses ini akan me-return data bertipe error
	err := json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("Tambah data gagal ! \n", err.Error())
		return
	}

	// jika tidak error, maka data ditambahkan kedalam slice users
	users = append(users, user)
	fmt.Println("User setelah ditambah dari frontend :", len(users))
	fmt.Println(users)
}