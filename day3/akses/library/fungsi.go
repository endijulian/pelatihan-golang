package library

import "fmt"

func Pesan(nama string) { //Jika nama function di awalai huruf Kapital (Public), jika sebaliknya (Private)

	fmt.Printf("Halo, Selamat datang %v", nama)
}

type Karyawan struct {

	//Huruf awal variabel juga mempengaruhi akses public atau private
	Nama   string
	Alamat string
}
