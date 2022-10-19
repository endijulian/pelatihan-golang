package main

import (
	"fmt"
	"strings"
)

type karyawan struct {
	namaPanjang string
}

func (data karyawan) pesan(kata string) {
	fmt.Println("Halo", data.namaPanjang)
}

func (data karyawan) namaDepan() string {
	return strings.Split(data.namaPanjang, "")[0]
}

func main() {
	var k1 = karyawan{"Endi Julian"}
	k1.pesan("Selamat Datang")

	fmt.Println("Nama Panggilan :", k1.namaDepan())
}
