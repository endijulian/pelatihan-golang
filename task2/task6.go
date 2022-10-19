package main

import "fmt"

func main() {

	var jenisKendaraan string
	var lamaParkir int

	println()
	fmt.Println("::: Park Team ::")
	println()

	fmt.Printf("Jenis Kendaraan :")
	fmt.Scanf("%v \n", &jenisKendaraan)

	fmt.Printf("Lama Parkir :")
	fmt.Scanf("%v \n", &lamaParkir)

	println("--------------------------------")

	fmt.Println("Kendaraan :", jenisKendaraan)
	// fmt.Printf("Biaya", biayaParkir(lamaParkir))
}

func tarifKendaraan(jenisKendaraan string) int {
	var tarif int
	if jenisKendaraan == "Motor" {
		tarif = 2000
	} else if jenisKendaraan == "Mobil" {
		tarif = 2000
	}

	return tarif
}

func biayaParkir(nilai int) {
	jumlahParameter := 0
	rata_rata := jumlahParameter * nilai

	fmt.Printf("Biaya : %v\n", rata_rata)
}
