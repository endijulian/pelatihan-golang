package main

import "fmt"

func main() {

	//Menambah dan membuat slice baru dengan append()
	var mobil = []string{"Honda", "Daihatsu", "Toyota"}
	var mobil2 = append(mobil, "Mitsubishi")

	fmt.Println(mobil)
	fmt.Println(mobil2)
	fmt.Println()

	mobil2[2] = "Datsun"

	fmt.Println(mobil)
	fmt.Println(mobil2)
	println()

	//Hubungan append() dengan refernce slice
	var buah = []string{"Buah", "Apel", "Salak", "Jambu"}
	var buahBaru = buah[0:3]

	// Nilai len(buahBaru) < cap(buahBaru)
	fmt.Println(len(buahBaru))
	fmt.Println(cap(buahBaru))
	fmt.Println()

	fmt.Println(buah)
	fmt.Println(cap(buahBaru))
	fmt.Println()

	//Pembuatan slice baru dengan append() dari slice buahBaru yang merupakan reference buah
	var buahBaru2 = append(buahBaru, "Pepaya")
	fmt.Println(buahBaru)
	fmt.Println(buahBaru2)

	buah[3] = "Anggur"
	fmt.Println(buah)
	fmt.Println(buahBaru2)

}
