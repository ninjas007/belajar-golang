package main

import (
	"fmt"
	// "time"
)

func main(){
	// name1 := "NooBee"
	// name2 := name1

	// name3 := "Beryl";
    // fmt.Println("Hello World")
	// contohVariable()
	// contohPercabangan()
	// contohArray()

	// contohSlice()
	// contohMap()
	// segitiga()
	// contohLoop()
	// contohCapdanLenAppend()
	// contohMapDelete()
	// contohFilter()
	// contohSliceOfMap()
	// contohStruct()
	// contohEmbedStruct()
	// contohSliceOfStruct()
	// contohFunction()
	// contohFunctionMultipleReturn()
	// contohMethod()
	contohPointer()

}

func contohPointer() {
	var num1 int = 5
    var num2 *int = &num1

    fmt.Println("===== Nilai Awal ===== ")
    fmt.Println(num1)
    fmt.Println(num2)
    fmt.Println(*num2)

    
    fmt.Println("===== Nilai Setelah num1 Diubah ===== ")
    num1 = 6
    fmt.Println(num1)
    fmt.Println(num2)
    fmt.Println(*num2)

    fmt.Println("===== Nilai Setelah num2 Diubah ===== ")
    *num2 = 10
    fmt.Println(num1)
    fmt.Println(num2)
    fmt.Println(*num2)
}

func contohVariable() {
	// # explisit
	// var name string
    // name = "NooBee"

    // var age int = 22
    // fmt.Printf("Halo, perkenalkan namaku %s dan aku berumur %d tahun\n", name, age)


	// # implisit dan deklarasi multilpe variable
	// name, age, isSingle := "NooBee", 22, true
    // fmt.Println(name, age, isSingle)


	// # pengenelanan tipe integer, harus sesuaikan data yang mau digunakan tidak melebihi dari nilainya
	// var age uint8 = 251 // maksimal uint8 adalah 255
    // fmt.Println("Umurku ", age)

	// var contohDecimal float32 = 1.66
    // fmt.Printf("%f\n", contohDecimal)
    // fmt.Printf("%.1f\n", contohDecimal)
    // fmt.Printf("%.3f\n", contohDecimal)

	// Latihan Variable
	/**
	Ubah kode dibawah ini menjadi lebih singkat
	*/

	// // sekarang
	// umurSaya := 20
	// umurKakak := 30
	// fmt.Printf("Umur saya sekarang adalah : %d\n", umurSaya)
	// fmt.Printf("Umur kakak sekarang adalah : %d\n", umurKakak)

	// // tahun depan
	// umurSaya++
	// umurKakak++

	// fmt.Printf("Umur saya tahun depan adalah : %d\n", umurSaya)
	// fmt.Printf("Umur kakak tahun depan adalah : %d\n", umurKakak)

	// // 5 tahun lagi
	// umurSaya+=5
	// umurKakak+=5

	// fmt.Printf("Umur saya 5 tahun lagi adalah : %d\n", umurSaya)
	// fmt.Printf("Umur kakak 5 tahun lagi adalah : %d\n", umurKakak)
}

func contohPercabangan() {
	// # scope variable dan percabangan
	// num := 2
	// text := "Test"
    // if num > 3 {
    //     text = "Lebih besar dari 1"
    // } else {
    //     text = "Lebih kecil atau sama dengan 1"
    // }

	// fmt.Println(text)

	// Latihan Percabangan
	/**
        Silahkan buat lakukan pengecekan apakah suatu variable bernilai ganjil atau genap.
        Contoh :
        num = 4 => genap
        num = 5 => ganjil

        Clue : 
        Menggunakan modulus, dan percabangan
	*/

	// num := 5

	// if num % 2 == 0 {
	// 	fmt.Printf("Bilangan %d adalah genap\n", num)
	// } else {
	// 	fmt.Printf("Bilangan %d adalah ganjil\n", num)
	// }


	// switch case
	// time := "am"

    // switch time {
    // case "pm" :
    //     fmt.Println("Malam")
    // case "am" :
    //     fmt.Println("Pagi")
    // default:
    //     fmt.Println("Maaf, waktunya salah !")
    // }

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("Hari ini libur!")
	// default:
	// 	fmt.Println("Hari ini kerja!")
	// }
}

func contohArray() {
	// # array

	// ... adalah tidak mendeskripsikan panjang datanya
	// fruits := [...]string{
    //     "Apple",
    //     "Banana",
    //     "Orange",
    // }

    // fmt.Println(fruits)

	// butuh paramter pertama sebagai tipe data
	// fruits := make([]string, 3)
    // fruits[0] = "Apple"
    // fruits[1] = "Banana"
    // fruits[2] = "Orange"

    // fmt.Println(fruits)


	// # multidimensional array
	// var fruits = [2][3]string{ 
    //     [3]string{"Apple", "Mango", "Banana"}, 
    //     [3]string{"Orange", "Grape","Avocado"},
    // }

    // fmt.Println(fruits)
    // fmt.Println(fruits[0])
    // fmt.Println(fruits[1])
}

func contohSlice() {

	// # slice
	fruits := []interface{}{1,"Banana", "Mango"} // ini slice
    animals := [3]string{"Cat", "Dog", "Elephant"} // ini array

    // cara menampilkannya sama
    fmt.Println(fruits[0])
    fmt.Println(animals[2])

	// slice reference ke array
	// animals := [4]string{"Cat", "Dog", "Chicken", "Bird"}
	// length := len(animals)

    // // legs4 := animals[2:3] 
    // legs2 := animals[2:4]

    // fmt.Println(length)
    // fmt.Println(legs2)

	// # slice type data reference
	// animals := [4]string{"Cat", "Dog", "Chicken", "Bird"}

    // anim1 := animals[0:3] 
    // anim2 := animals[2:4]

    // fmt.Println(animals)    // [Cat, Dog, Chicken, Bird]
    // fmt.Println(anim1)      // [Cat, Dog, Chicken]
    // fmt.Println(anim2)      // [Chicken, Bird] 
	// fmt.Println("\n")


    // anim1[1] = "Cow"
    // fmt.Println(animals)    // [Cat, Cow, Chicken, Bird]
    // fmt.Println(anim1)      // [Cat, Cow, Chicken]
    // fmt.Println(anim2)      // [Chicken, Bird]
	// fmt.Println("\n")


    // anim2[0] = "Pinguin"
    // fmt.Println(animals)    // [Cat, Dog, Pinguin, Bird]
    // fmt.Println(anim1)      // [Cat, Dog, Pinguin]
    // fmt.Println(anim2)      // [Chicken, Pinguin]
	

	// # LATIHAN
	// // Diberikan sebuah array seperti berikut :
    // animals := [...]string{"Cat", "Dog", "Pinguin", "Chicken", "Snake"}

    // Lalu, lengkapi variable variable berikut sesuai dengan expected-nya :
    // mammals := animals[0:2]     // expected : {Cat, Dog}
    // notMammals := animals[2:5]  // expected : {Pinguin, Chicken, Snake}
    // haveLegs := animals[0:4]    // expeccted : {Cat, Dog, Pinguin, Chicken}

	// fmt.Println(mammals)    // expected : {Cat, Cow}
    // fmt.Println(notMammals) // expected : {Bird, Chicken, Snake}
    // fmt.Println(haveLegs)  

	// fmt.Println("\n\n\n\n")

    // Setelah itu, lakukan hal berikut :
    // 1. Ubah value Dog menjadi Cow
    // 2. Ubah value Pinguin menjadi Bird

    // dimulai dari sini
	// animals[1] = "Cow"
	// animals[2] = "Bird"



    // // // berakhir disini

    // // Saat di print, harusnya hasilnya sesuai dengan expected
    // fmt.Println(mammals)    // expected : {Cat, Cow}
    // fmt.Println(notMammals) // expected : {Bird, Chicken, Snake}
    // fmt.Println(haveLegs)   // expected : {Cat, Cow, Bird, Chicken}
}

func contohMap() {
	person := make(map[string]string) // inisialisasi map kosong
    person2 := map[string]string{
        "name" : "NooBee", 
        "job" : "Programmer",
    } // inisialisasi map dengan langsung mendefinisikan key dan value

    computer := map[string]int{
        "sold" : 14,
        "stock" : 20,
    } // inisialisasi map dengan key dan value yang berbeda tipe data

    fmt.Println(person) // expect: map[]
    fmt.Println(person2) // expect: map[name:NooBee job:Programmer]
    fmt.Println(computer) // expect: map[stock:20 sold:14]
    fmt.Println("")

    person["name"] = "Reyhan"
    person["job"] = "Designer"

    person2["address"] = "Pekanbaru"

    computer["created_year"] = 2019

    fmt.Println(person)
    fmt.Println(person2)
    fmt.Println(computer)

    for key, value := range person2 {
        fmt.Printf("key : %s, value : %s dan nama adalah  %s \n", key, value, person["name"])
    }
}

func segitiga() {

	for i := 5; i > 0; i-- {

		for j := 0; j < i; j++ {

			fmt.Print("*")

		}

		fmt.Println()
	}
}

func contohLoop() {
	// for {
	// 	fmt.Println("ini ga bakal berhenti")
	// }

	// Latihan 1
    /** 
        Silahkan buat bentuk seperti ini :
        *****
        ****
        ***
        **
        *
    */
	for i := 10; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

    // Latihan 2
    /** 
        Silahkan buat bentuk seperti ini :
        *
        **
        ***
        ****
        *****
    */

	for i := 0; i <= 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

    // Latihan 3
    /**
        Silahkan buat bentuk Noo Bee (panjang : 50) dengan kondisi berikut :
        Jika habis dibagi 3 : Noo
        Jika habis dibagi 5 : Bee
        Jika habis dibagi 3 dan 5 : Noo Bee

        Contoh :
        1, 2, Noo, 4, Bee, Noo, 7, 8, Noo, Bee, 11, Noo, 13, 14, NooBee, 16, 17, Noo, 19, Bee, ... n
    */

	for i := 1; i <= 50; i++ {
		if i % 3 == 0 {
			
			if i % 5 == 0 {
				fmt.Print("Noo Bee, ")
			}

			fmt.Print("Noo, ")
		}
		
		if i % 5 == 0 {
			fmt.Print("Bee, ")
		}
	}
}

func contohCapdanLenAppend() {
	animals := []string{"Cat", "Dog", "Cow"}
    animSmall := animals[0:2] // expected : Cat, Dog

    anim1 := append(animals, "Sheep") // expected: Cat, Dog, Cow, Sheep
    anim2 := append(animSmall, "Ant") // expected: Cat, Dog, Ant

    fmt.Println("====[ANIMALS]====")
    fmt.Println("Capacity :",cap(animals))
    fmt.Println("Len \t :",len(animals))
    fmt.Println("Data \t :", animals)
    fmt.Println("")

    fmt.Println("====[ANIMALS SMALL]====")
    fmt.Println("Capacity :",cap(animSmall))
    fmt.Println("Len \t :",len(animSmall))
    fmt.Println("Data \t :", animSmall)
    fmt.Println("")


    fmt.Println("====[ANIM1]====")
    fmt.Println("Capacity :",cap(anim1))
    fmt.Println("Len \t :",len(anim1))
    fmt.Println("Data \t :", anim1)
    fmt.Println("")

    fmt.Println("====[ANIM2]====")
    fmt.Println("Capacity :",cap(anim2))
    fmt.Println("Len \t :",len(anim2))
    fmt.Println("Data \t :", anim2)
    fmt.Println("")
}

func contohMapDelete() {
	nilai := map[string]int{
        "Agama" : 80,
        "Matematika" : 90,
        "Olahraga" : 70,
        "Design" : 80,
    }

    fmt.Println("===== Nilai Awal ======")
    fmt.Println(len(nilai))
    fmt.Println(nilai)
    fmt.Println("\n\n===== Nilai Akhir ======")
    delete(nilai, "Design") // hapus nilai disini
    fmt.Println(len(nilai))
    fmt.Println(nilai)
}

func contohFilter() {
	nilai := map[string]int{
        "Agama" : 80,
        "Matematika" : 90,
        "Olahraga" : 70,
        "Design" : 80,
    }

    key := "Agama"
    val, isExist := nilai[key]

    if isExist {
        fmt.Println(key,"is exist with value", val)
    } else {
        fmt.Println(key,"is not exist")

    }
}

func contohSliceOfMap() {
	students := []map[string]string{
        map[string]string{"name":"NooBee", "major":"Teknik Komputer"},
        map[string]string{"name":"Jovie", "major":"Sistem Informasi"},
        map[string]string{"name":"Reyhan", "major":"Teknik Informatika"},
    }

    for _, student := range students {
        fmt.Println(student["name"], "berada di jurusan", student["major"])
    }
}

func contohStruct() {
	type Fruit struct {
		Name string
		Weight int
	}

	var fruit1 Fruit
    fruit1.Name = "Apple"
    fruit1.Weight = 1

    fruit2 := Fruit{
        Name : "Mango",
        Weight : 2,
    }

    var fruit3 = Fruit{
        Name : "Banana",
        Weight : 1,
    }

    fruit4 := Fruit{"Lemon", 4}

    fmt.Println(fruit1)
    fmt.Println(fruit2)
    fmt.Println(fruit3)
    fmt.Println(fruit4)
}

func contohEmbedStruct() {

	type Parent struct {
		Nama string
		Umur int
	}
	
	type Student struct {
		OrangTua Parent
		Nama string
		Umur int
		Kelas string
	}

	parent1 := Parent{
        Nama : "Andi",
        Umur : 50,
    }

    student1 := Student{
        OrangTua : parent1,
        Nama : "Budi",
        Umur : 11,
        Kelas : "6A",
    }

    student2 := Student {
        OrangTua : Parent{
            Nama: "Jojo",
            Umur: 51,
        },
        Nama: "Bilqis",
        Umur: 10,
        Kelas: "5B",
    }

    fmt.Println(student1)
    fmt.Println(student2)

    fmt.Printf("Orang Tua %s Bernama %s\n", student1.Nama, student1.OrangTua.Nama)
    fmt.Printf("Orang Tua %s Bernama %s\n", student2.Nama, student2.OrangTua.Nama)
}

func contohSliceOfStruct() {
	type Employee struct{
		Name string
		Departement string
		test int
	}
	
	type Employees []Employee

	var employees Employees
    fmt.Println(employees)

    // employees[0].Name = "Employee A" => ini akan error dikarenakan saat deklarasi awal, employees tidak mempunyai element sama sekali (slice kosong)

    var emp1 = Employee{Name:"Emp1", Departement: "Tech"}
    employees = append(employees, emp1)

    fmt.Println(employees)

    employees = append(employees, Employee{Name:"Emp2", Departement:"Finance"})
    fmt.Println(employees)
}

func contohFunction() {
	// var names = []string{"Rey", "Jo", "NooBee"}
    // var ages = []int{10, 13, 20}
    // for _, name := range names {
    //     print(name)
    // }

    // for _, age := range ages {
    //     print(age)
    // }

	// fmt.Println()
	// fmt.Println("Contoh Function return")
	// fmt.Println(hitung(10, 20, "+"))

	// Latihan
	fmt.Println("Latihan")

	var car = make(map[string]string)
    car["name"] = "BWM"
    car["color"] = "Black"

    // buat 2 buah fungsi :
    // 1 => fungsi yang mengembalikan sebuah string
    // pada fungsi ini terjadi pengolahan kata sehingga menghasilkan kata : Mobil BMW berwarna Black
	msg := stringCar(car["name"], car["color"])

    // 2 => fungsi yang menampilkan hasil dari kembalian string
    // fungsi ini hanya bertugas untuk menampilkan kata
	print(msg);

    // alur
    // simpan hasil dari return function kedalam sebuah variable message
    // tampilkan hasil dari variable message


    // output => Mobil BMW berwarna Black
}

// make function argument type data interface
func print(arg interface{}) {
	fmt.Println(arg)
}

func stringCar(name string, color string) string {
	return "Mobil " + name + " berwarna " + color;
}

func hitung(a int, b int, operator string) int {

	if operator == "+" {
		return a + b
	} else if (operator == "-") {
		return a - b
	} else if (operator == "*") {
		return a * b
	} else if (operator == "/") {
		return a / b
	} else {
		return 0;
	}
}

func contohFunctionMultipleReturn() {
	var sisi float32 = 4

    luas, keliling := calculate(sisi)
    fmt.Println("Luas dan Keliling dari persegi dengan panjang sisi",sisi)
    fmt.Println("Luas:", luas)
    fmt.Println("Keliling:", keliling)

	// function closure
	calculate := func (num1, num2 int) int {
        return num1 + num2
    }

    fmt.Println(calculate(1,3))

	// function variadic
	fmt.Println(sum(1,8,29,3,4,5,2,9,7))
}

func calculate(sisi float32)(float32, float32){
    luas := sisi * sisi
    keliling := sisi * 4
    return luas, keliling
}

func sum(num ...int) (total int) {
    fmt.Println(num)

    for _, n := range num {
        total += n
    }

    return total
}

// contoh method
type Vehicle struct {
	Name string
	Color string
}

func (v Vehicle) GetName() string {
	return v.Name
}

func (v Vehicle) SayHello(){
	fmt.Println("Hello dari", v.Name, v.Color)
}

func contohMethod() {
	car := Vehicle{
		Name : "Civic",
		Color : "White",
	}

	bike := Vehicle{
		Name : "Ducati",
		Color : "Black",
	}

	nameCar := car.GetName()
	nameBike := bike.GetName()
	fmt.Println("Nama kendaraannya adalah:", nameCar, "dan", nameBike)
	car.SayHello()
}

func sortPhp() {
// 	<?php
// // Online PHP compiler to run PHP program online
// // Print "Hello World!" message\

// $arr = [2, 3, 7, 9, 5, 1];
// $length = count($arr);
// $curr = [];
// for ($i = 0; $i < $length; $i++) {
//     for ($j = $i + 1; $j < $length; $j++) {
//         if ($arr[$i] > $arr[$j]) {
//             echo PHP_EOL;
//             echo PHP_EOL;
//             echo "++++++++++++++++";
//             echo "ubah angka $arr[$i] di index $i ke $arr[$j] di index $j";
//             echo PHP_EOL;
            
//             $temp = $arr[$i];
            
//             $arr[$i] = $arr[$j];
            
//             $arr[$j] = $temp;
            
//             echo PHP_EOL;
//             echo "Keadaan Sekarang";
//             echo PHP_EOL;
//             echo join(',', $arr);
//             echo PHP_EOL;
//         }
//     }
// }

// ?>
}
