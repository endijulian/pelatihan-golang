package main

import "fmt"

func main() {
	cetak("Bayu", "Selamat Datang")

	//Pemanggilan return
	fmt.Println(hitungJumlah(5, 4))

	var hasilHitung = hitungJumlah(9, 20)
	fmt.Println(hasilHitung)

	//Pemanggilan mulitple return
	nilaiLuas, nilaiKeliling, namaBentuk := persegi(6)
	fmt.Printf("Bangun datar %v, Luas = %v, Keliling = %v\n", namaBentuk, nilaiLuas, nilaiKeliling)
}

//Pembuatan function non-return
func cetak(nama string, pesan string) {

	//Pemanggilan fungsi non-return
	fmt.Printf("%v, %v di Rumah Coiding", nama, pesan)
	println()

}

//Pembuatan function return
func hitungJumlah(nilai1 int, nilai2 int) int {
	hasil := nilai1 + nilai2
	return hasil
}

//Penulisan beberapa nama parameter yang tipe datanya sama
func luasPersegiPanjang(panjang, lebar int) uint8 {
	luas := panjang * lebar

	return uint8(luas)
}

//Function multiple return
func persegi(sisi int) (int, int, string) {
	luas := sisi * sisi
	keliling := sisi * 4
	namaBentuk := "Persegi"

	return luas, keliling, namaBentuk
}

//Function predefine return / definisi nama return di awal
func bangunPersegi(sisi int) (luas int, keliling int, namaBentuk string) {
	luas = sisi * sisi
	keliling = sisi * 4
	namaBentuk = "Persegi"

	return
}
