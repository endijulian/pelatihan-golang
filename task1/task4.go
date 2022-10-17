package main

import "fmt"

func main() {
	var pemebelian int = 0
	layanan := ""

	fmt.Printf("Jumlah Pembelian / Lembar :")
	fmt.Scan(&pemebelian)

	fmt.Printf("Jenis Layanan :")
	fmt.Scan(&layanan)

	// fmt.Println("Total :")
	// fmt.Scan()

	// fmt.Println("Diskon :")
	// fmt.Scan()

	// fmt.Println("Jumlah yang harus dibayar :")
	// fmt.Scan()

	//Output
	fmt.Printf("Jumlah Pembelian / Lembar: %d \n", pemebelian)
	fmt.Println("Layanan :", layanan)
	if layanan == "Warna" {
		fmt.Printf("Total :%d \n", (pemebelian * 1000))
		// harga := pemebelian * 1000
	} else if layanan == "Hitam" {
		fmt.Printf("Total :%d \n", (pemebelian * 125))
		// harga := pemebelian * 125
	}

}
