package main

import "fmt"

func main() {
	var product = [2][5]string{
		{"Keyboard", "Mouse", "LCD", "LED", "Headset"},
		{"Ram", "HDD", "Motherboard", "VGA", "SSD"},
	}

	var cari, ganti string
	// var change string

	productBaru := product[0:2]

	fmt.Println("Data Product sebelum di ubah :")
	fmt.Println("Product :", product)
	fmt.Println("Product Baru:", productBaru)

	fmt.Printf("Masukan nama barang yang ingin diganti : ")
	fmt.Scanf("%v\n", &cari)

	for key1, index := range product {
		for key2, barang := range index {
			if cari == barang {
				fmt.Printf("Nama Product ditemukan, ganti dengan nama :")
				fmt.Scanf("%v \n", &ganti)
				productBaru[key1][key2] = ganti
			}
		}
	}

	println()
	println()

	fmt.Println("Data Produk Setelah diubah :")
	fmt.Println("Product :", product)
	fmt.Println("Product Baru:", productBaru)
}
