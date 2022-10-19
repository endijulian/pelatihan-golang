package main

import "fmt"

type siswa struct {
	nama   string
	gender string
	usia   int
}

func inputSiswa(nama string, gender string, usia int) siswa {
	input := struct {
		nama   string
		gender string
		usia   int
	}{
		nama:   nama,
		gender: gender,
		usia:   usia,
	}

	return input
}

func tampilSiswa(data siswa) {
	fmt.Printf("Nama : %v\n", data.nama)
	fmt.Printf("Gender : %v\n", data.gender)
	fmt.Printf("Usia : %v\n", data.usia)
}

func main() {
	//Input data menggunakan fungsi
	var siswa1 = inputSiswa("Hari", "Laki-Laki", 25)
	var siswa2 = inputSiswa("Jono", "Laki-Laki", 28)

	tampilSiswa(siswa1)
	tampilSiswa(siswa2)

	// fmt.Println(tampilSiswa(siswa1))
}
