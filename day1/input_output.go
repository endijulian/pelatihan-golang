package main

import "fmt"

func main() {

	var nama string
	var alamat string

	//Input menggunakan scan
	fmt.Printf("Masukan Nama : ")
	fmt.Scanf("%s\n", &nama)

	//Input menggunakan scan scanf
	fmt.Printf("Masukan Alamat : ")
	fmt.Scanf("%s\n", &alamat)

	//Menampilkan hasil input
	fmt.Println()
	fmt.Printf("Hasil Input : %s", nama)

	fmt.Println()
	fmt.Printf("Alamat anda di : %s", alamat)
}
