package main

import "fmt"

func main() {

	//Inisialisasi map dalam slice

	var karyawan = []map[string]string{
		{"nama": "Bayu Akbar", "alamat": "Depok"},
		{"nama": "Dani", "alamat": "Jakarta"},
		{"nama": "Isnawati", "alamat": "Bogor"},
	}

	for _, v := range karyawan {
		// fmt.Println(v["nama"], v["alamat"])
		// fmt.Printf("Nama :%v \nAlamat : %v \n", v["nama"], v["alamat"])
		fmt.Printf("Nama :%v \n", v["nama"])
		fmt.Printf("Alamat :%v \n", v["alamat"])
	}
}
