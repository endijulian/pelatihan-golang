package main

import "fmt"

func main() {
	var character = [...]string{"Hero", "Monster", "Goblin"}

	//Menampilkan seluruh nilai array dengan FOR
	jumlahData := len(character)
	println(jumlahData)

	for i := 0; i < jumlahData; i++ {
		fmt.Printf("Elemen ke -%v, Nilai = %v \n", i, character[i])
	}

	println()
	//Menampilkan seluruh nilai array dengan FOR RANGE
	for index, nilai := range character {
		fmt.Printf("Elemen ke - %v, Nilai = %v\n", index, nilai)
	}

	println()
	//Menggunakan underscore (_) untuk menaggulangi error untuk variabel yang tidak digunakan
	for _, nilai := range character {
		fmt.Printf("Nilai = %v\n", nilai)
	}
}
