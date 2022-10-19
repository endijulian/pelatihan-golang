package main

import "fmt"

func main() {
	var angkaA int = 5
	var angkaB *int = &angkaA //Inisialisasi variabel pointer angkaB

	fmt.Println("Nilia Angka A :", angkaA)
	fmt.Println("Alamat Memori Angka A :", &angkaA) //Operator & berfungsi untuk mendapatkan alamat memori variabel biasa
	fmt.Println("Nilia Angka B :", *angkaB)         //Operator * untuk menampilkan nilai asli variabel referensi dari variabel pointer
	fmt.Println("Alamat Memori Angka B :", angkaB)

	fmt.Println()
	fmt.Println()

	*angkaB = 10 //Ubah nilai referensi variabel pointer
	println()

	//Menampilkan nilai setelah di ubah
	fmt.Println("Nilia Angka A :", angkaA)
	fmt.Println("Alamat Memori Angka A :", &angkaA)
	fmt.Println("Nilia Angka B :", *angkaB)
	fmt.Println("Alamat Memori Angka B :", angkaB)
}
