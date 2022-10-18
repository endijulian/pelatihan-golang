package main

import "fmt"

func main() {

	//Deklarasi MAP tanpa mengisi nilai
	var biodata map[string]string
	biodata = map[string]string{}

	//Isi nilai map yang sudah di buat
	biodata["nama"] = "Bayu"
	biodata["alamat"] = "depok"

	fmt.Printf("Nama anda %v \n", biodata["nama"])
	fmt.Printf("Alamat anda %v \n", biodata["alamat"])

	//Inisialisasi / deklarasi variabel dengan langsung mengusu nilai
	var mobil = map[string]string{"merk": "Daihatsu", "tipe": "Rocky"} //Secara horizontal
	var motor = map[string]string{
		"merk":  "Honda",
		"tipe":  "Beat",
		"jenis": "Sepeda Motor",
	}

	fmt.Println(mobil)
	fmt.Println(motor)

	//Menampilkan seluruh nilai MAP motor dengan FOR
	for key, value := range motor {
		fmt.Printf("%v  =  %v \n", key, value)
	}

	//Hapus nilai map menggunakan fungsi delete()
	delete(motor, "jenis")
	for key, value := range motor {
		fmt.Printf("%v  =  %v \n", key, value)
	}

}
