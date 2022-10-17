package main

import "fmt"

func main() {

	var hari string
	layanan := 2
	var harga1 int = 45000
	var harga2 int = 5000

	fmt.Printf("Masukan Hari Servie :")
	fmt.Scanf("%s\n", &hari)

	fmt.Printf("Layanan :\n")
	fmt.Println("1. Bensin")
	fmt.Println("2. Service")
	fmt.Println("3. Bensin + Service")

	fmt.Printf("Pilih Layanan (1/2/3) :")
	fmt.Scanf("%d\n", &layanan)

	fmt.Printf("Hari Service : %s \n", hari)
	if layanan == 1 {
		fmt.Println("Layanan :", "Bensin")
	} else if layanan == 2 {
		fmt.Println("Layanan :", "Service")
	}

	if layanan == 1 {
		fmt.Println("Biaya :", 10000)
	} else if layanan == 2 {
		fmt.Println("Biaya :", harga1)
	}

	if hari == "Rabu" {
		fmt.Println("Diskon :", harga2)
	} else {
		fmt.Println("Diskon :")
	}

	fmt.Println("Total Bayar :", (harga1 - harga2))

}
