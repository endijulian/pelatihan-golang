package main

import "fmt"

//Pembuatan Interface
type hitung interface {
	luas() float64
	panjang() float64
}

//Struct dan method lingkaran
type lingkaran struct {
	diameter float64
}

func (nilai lingkaran) jarijari() float64 {
	return nilai.diameter / 2
}

func (nilai lingkaran) luas() float64 {
	return 3.14 * (nilai.jarijari() * nilai.jarijari())
}

func (nilai lingkaran) keliling() float64 {
	return 3.14 * nilai.diameter
}

//Struct dan method persegi
type persegi struct {
	sisi float64
}

func (nilai persegi) luas() float64 {
	return nilai.sisi * nilai.sisi
}

func (nilai persegi) keliling() float64 {
	return nilai.sisi * 4
}

func main() {
	var bangunDatar hitung

	// bangunDatar = persegi{10.0}
	fmt.Println("========== Persegi =========")
	fmt.Println("Luas			:", bangunDatar.luas())
	// fmt.Println("Keliling		:", bangunDatar.keliling())

	// bangunDatar = lingkaran{14.0}
	fmt.Println("========== Lingkaran =========")
	fmt.Println("Luas			:", bangunDatar.luas())
	// fmt.Println("Keliling		:", bangunDatar.keliling())
	// fmt.Println("Jari-Jari		:", bangunDatar.(lingkaran).jarijari())

}
