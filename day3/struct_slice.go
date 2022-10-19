package main

import "fmt"

type Biodata struct {
	nama   string
	alamat string
	gender string
}

func main() {

	//Inisialisasi langsung nilai semua data
	var semuaData = []Biodata{
		{nama: "Endi", alamat: "Tangerang", gender: "Anonym"},
		{nama: "Julian", alamat: "Bekasi", gender: "Perempuan"},
		{nama: "Array", alamat: "Jakarta", gender: "Laki-Laki"},
	}

	//Inisialisasi slice struct anonymouse
	var semuaKaryawan = []struct {
		divisi  string
		Biodata // embeded struct dari luar fungsi main ke dalam anonymouse struct
	}{
		{divisi: "IT", Biodata: Biodata{nama: "Endi", alamat: "Tangerang", gender: "Anonym"}},
		{divisi: "Keuangan", Biodata: Biodata{nama: "Julian", alamat: "Bekasi", gender: "Perempuan"}},
		{divisi: "HR", Biodata: Biodata{nama: "Array", alamat: "Jakarta", gender: "Laki-Laki"}},
	}

	fmt.Printf("Semua karyawan: %v\n", semuaKaryawan)

	fmt.Println()
	fmt.Println()
	//Menampilkan seluruh nilai menggunakan for
	for _, v := range semuaData {
		fmt.Printf("Nama : %v\nAlamat : %v\nJenis Kelamin : %v\n", v.nama, v.alamat, v.gender)
	}
}
