package main

import "fmt"

func main() {

	//Perulanga tipe pertama
	for i := 0; i < 5; i++ {
		println("Angka :", i)
	}
	println()
	println()

	//Perulangan tipe ke 2
	x := 0
	for x < 5 {
		println("Nilai :", x)
		x++
	}
	println()
	println()

	//Break dan continue dalam for
	for x := 0; x <= 9; x++ {
		if x < 8 {
			continue
		}

		if x > 8 {
			break
		}

		println("Cetak", x)
	}

	//Nested for (Bersarang)
	for w := 0; w < 7; w++ {
		fmt.Println("Menaik", w)
		for d := 10; d > 0; d-- {
			fmt.Println("Menurun :", d)
		}
	}
}
