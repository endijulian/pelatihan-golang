package main

import "fmt"

func main() {
	var panjang int
	var lebar int
	var tinggi int

	fmt.Printf("Masukan panjang :")
	fmt.Scanf("%d\n", &panjang)

	fmt.Printf("Masukan lebar :")
	fmt.Scanf("%d\n", &lebar)

	fmt.Printf("Masukan panjang :")
	fmt.Scanf("%d\n", &tinggi)

	fmt.Println("Hasil :", (panjang * lebar * tinggi))
}
