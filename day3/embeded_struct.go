package main

import "fmt"

type Biodata struct {
	nama   string
	alamat string
	usia   int
}

//Embedded Struct
type Karyawan struct {
	divisi  string
	nama    string
	Biodata //Struct yang dugunakan oleh struct lain
}

func main() {

	//Penggunaan Struct embedded pada object baru
	var k1 = Karyawan{}
	k1.nama = "Julian"
	k1.Biodata.nama = "Randy Nugroho"
	k1.alamat = "Jakarta"
	k1.usia = 24
	k1.divisi = "IS"

	fmt.Println("Nama Karyawan :", k1.nama)
	fmt.Println("Nama Struct Biodata :", k1.Biodata.nama)
	fmt.Println("Alamat Karyawan :", k1.alamat)
	fmt.Println("Usia Karyawan :", k1.usia)
	fmt.Println("Divisi Karyawan :", k1.divisi)
}
