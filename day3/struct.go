package main

import "fmt"

//

func main() {
	//Penggunaan struct didalam main

	//Deklarasi
	var kar1 Karyawan

	//Pengisian nilai property object kar1
	kar1.nama = "Endi Julian"
	kar1.status = true

	fmt.Printf("Nama Karyawan %v\n", kar1.nama)
	fmt.Printf("Status Karyawan %v\n", kar1.status)
	println()
	println()

	//Inisialisasi / pengisian nilai langsung ke object struct baru
	var kar2 = Karyawan{"Doni", true}
	kar3 := Karyawan{status: true, nama: "Andini"}

	fmt.Printf("KAR 2 : %v\n", kar2)
	fmt.Printf("KAR 3 : %v\n", kar3)
	println()
	println()

	//Variabel object pointer / membuat object struct baru berdasarkan object struct yang sudah ada
	var pointerKar2 *Karyawan = &kar2
	fmt.Printf("Pointer %v \n", pointerKar2.nama)
	fmt.Printf("Pointer %v", pointerKar2.status)
	println()
	println()

	pointerKar2.nama = "Bagas"
	fmt.Println("Nama Kar 2 :", kar2.nama)
	fmt.Println("Nama Pointer kar 2 :", pointerKar2.nama)

}

//Pembuatan Struct
type Karyawan struct {
	nama   string
	status bool
}
