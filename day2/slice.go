package main

import "fmt"

func main() {

	//Dekralasi SLICE
	var buah = []string{"Pisang", "Apel", "Semangka", "Melon"}

	//Akses / Menampilkan nilai slice
	fmt.Println(buah)
	fmt.Println(buah[0])
	fmt.Println(buah[2])

	//Mengubah nilai didalam SLICE
	buah[0] = "Durian"
	buah[3] = "Anggur"

	fmt.Println(buah)

	//Mengetahui jumlah elemen
	fmt.Println("Jumlah elemen buah", len(buah))

}
