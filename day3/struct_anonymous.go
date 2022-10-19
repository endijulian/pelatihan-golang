package main

import "fmt"

type Biodata struct {
	nama   string
	alamat string
}

func main() {

	//Membuat / deklarasi struct anonymouse dan langsung di inisialisasi ke variabel
	var karyawanJoni = struct {
		divisi  string
		Biodata //Embed struct dari luar fungsi main
	}{}

	//Penggunaan anonymouse struct
	karyawanJoni.nama = "Joni"
	karyawanJoni.alamat = "Depok"

	fmt.Printf("Karyawan Joni Nama : %v\n", karyawanJoni.nama)
	fmt.Printf("Karyawan Joni Alamat : %v\n", karyawanJoni.alamat)
}
