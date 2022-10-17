package main

import "fmt"

func main() {

	//Dekralasi variabel dengan menentukan tipe data
	var nama string = "Endi"
	var nilai int = 90
	var alamat string = "depok"

	//Konstanta
	const pi = 3.14
	const produk string = "Kopi"

	// %s untuk menampilkan tipe data string
	// %d untuk menampilkan tipe data int

	fmt.Printf("Halo %s Nilai Kamu : %d dan alamat rumah mu %s \n", nama, nilai, alamat)

	//Menggunakan println
	// fmt.Println("Hallo", nama)

	//Dekralasi variabel tanpa tipe data
	var namaKaryawan = "Endi Julian"
	alamatKaryawan := "Tangerang"
	usia := 24

	fmt.Printf("Nama %s, %s Usia %d \n", namaKaryawan, alamatKaryawan, usia)

	// Dekaralasi multiple variable

	// tipe1
	var nama1, nama2, nama3 string
	nama1, nama2, nama3 = "Bayu", "Akbar", "Saputro"

	// tipe 2
	var nilai1, nilai2 int = 90, 100

	//Tipe 3
	alamat1, usia1 := "Tangerang", 24

	fmt.Printf("%s %s %s %s %d %d %d \n\n", nama1, nama2, nama3, alamat1, nilai1, nilai2, usia1)
	fmt.Println(pi, produk)

}
