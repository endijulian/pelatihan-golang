package main

import "fmt"

func main() {
	var product = [2][5]string{
		{"Keyboard", "Mouse", "LCD", "LED", "Headset"},
		{"Ram", "HDD", "Motherboard", "VGA", "SSD"},
	}

	var cari string
	fmt.Printf("Cari Produk: ")
	fmt.Scanf("%v", &cari)

	//Menggunakan for range
	for i, index := range product {
		for j, barang := range index {
			if cari == barang {
				fmt.Printf("Product %v ada pada index ke [%v][%v] \n", barang, i, j)
			}
		}
	}
}
