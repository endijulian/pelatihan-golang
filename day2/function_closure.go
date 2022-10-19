package main

import "fmt"

func main() {

	//Function closure / anonymouse
	tampilPesan := func(nama string, pesan string) string {
		tampil := fmt.Sprintf("%v di Rumah Coding %v, Kita akan belajar golang\n", pesan, nama)
		return tampil
	}

	hasil := tampilPesan("Bayu", "Welcome")
	println(hasil)

	fmt.Println(tampilPesan("Doni", "Selamat datang"))

	//Closure IIFE
	nilai := []int{100, 90, 89, 88, 98, 55}
	var rataRata = func(data []int) int {
		jumlahNilai := 0
		for _, v := range data {
			fmt.Println(v)
			jumlahNilai += v
		}

		rata_rata := jumlahNilai / len(data)
		return rata_rata
	}(nilai)

	fmt.Printf("Seluruh nilai andi : %v\n", nilai)
	fmt.Printf("Siswa andi dengan nilai rata-rata : %v\n", rataRata)
	println()
	println()

	///Pemanggilan function return closure
	max := 70
	nilai1 := []int{100, 90, 80, 70, 60, 50}
	var jumlahFilter, nilaiFilter = filterMax(max, nilai1...)
	var hasilFilter = nilaiFilter()

	fmt.Println("Jumlah data filter :", jumlahFilter)
	fmt.Println("Nilia rata-rata setelah di filter :", hasilFilter)
}

func filterMax(max int, nilai ...int) (int, func() []int) {
	var hasil []int
	for _, v := range nilai {
		if v <= max {
			hasil = append(hasil, v)
		}
	}

	return len(hasil), func() []int { return hasil }
}
