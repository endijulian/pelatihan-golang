package main

import "fmt"

type Biodata struct {
	nama     string
	alamat   string
	Karyawan struct {
		divisi string
	}
}

func main() {

	var karyawan1 = Biodata{}
	karyawan1.nama = "Joni"
	karyawan1.alamat = "Depok"
	karyawan1.Karyawan.divisi = "IT"

	fmt.Println("Nama Karyawan", karyawan1.nama)
	fmt.Println("Alamat Karyawan", karyawan1.alamat)
	fmt.Println("Divisi Karyawan", karyawan1.Karyawan.divisi)

}
