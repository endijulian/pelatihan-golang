package main

import "fmt"

func main() {

	//Penggunaan function variadic dengan mengirim nilai langsung
	hasilRata := rataRata(3, 5, 7, 3, 4, 5, 3, 5)
	fmt.Println("Nilai rata-rata :", hasilRata)

	//Penggunaan function variadic dengan mengirim slice
	var sliceNilai = []int{3, 4, 5, 6, 9, 5, 3, 5, 7, 2}
	hasilRata1 := rataRata(sliceNilai...)
	fmt.Println("Nilai hasil rata-rata 1: ", hasilRata1)

	//Kirim argumen kombinasi variadic dan parameter biasa
	nilaiSiswa("Andi", 90, 100, 54, 67, 89, 90)

	var sliceNilaiSiswa = []int{100, 55, 90, 89, 78, 96}
	nilaiSiswa("Jono", sliceNilaiSiswa...)

}

//Pembutan function variadic
func rataRata(nilai ...int) int { //titik 3 ... menandakan parameter variadic

	jumlahParameter := 0
	for _, v := range nilai {
		jumlahParameter += v
	}

	rata_rata := jumlahParameter / len(nilai)
	return rata_rata
}

func nilaiSiswa(nama string, nilai ...int) {
	jumlahParameter := 0
	for _, v := range nilai {
		jumlahParameter += v
	}

	rata_rata := jumlahParameter / len(nilai)
	fmt.Printf("Siswa bernama %v dengan rata-rata nilai : %v\n", nama, rata_rata)
}
