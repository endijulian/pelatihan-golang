package main

import "fmt"

func main() {
	//Dekralasi varibel array
	var nama [3]string
	var nilai [3]int
	var bolean [2]bool

	//Input nilai variabel array
	nama[0] = "John"
	nama[1] = "Ikuwail"
	nama[2] = "Wickhle"

	nilai[0] = 80

	bolean[0] = true

	fmt.Println(nama)
	fmt.Println(nilai)
	fmt.Println(bolean)

	//Inisialisasi / isi data saat dekralasi variabel array

	var namaKaryawan = [3]string{"John", "Endi", "Trish"}
	var usia = [3]int{
		30,
		20,
		29,
	}

	alamat := [2]string{"Depok", "Bogor"}

	println()
	fmt.Println(namaKaryawan)
	fmt.Println(usia)
	fmt.Println(alamat)

	println()
	println()
	//Dekralasi variabel array tanpa jumlah elemen
	var product = [...]string{"keyboard", "laptop", "LCD"}
	var product2 = [4]string{"keyboard", "laptop", "LCD"}

	//Akses atau menampilkan variabel array
	fmt.Println(product[1])
	fmt.Println(product2[2])
	fmt.Println(product2[3])
}
