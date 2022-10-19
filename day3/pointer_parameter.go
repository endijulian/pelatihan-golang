package main

import "fmt"

func main() {
	var nilai int = 8
	fmt.Println("Nilai sebelum diubah :", nilai)
	ubahNilai(&nilai, 15)
	fmt.Println("Nilai setelah di ubah :", nilai)
}

//Pembuatan function dengan parameter pointer
func ubahNilai(ubah *int, nilai int) {
	*ubah = nilai
}
