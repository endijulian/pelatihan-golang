package main

import "fmt"

func main() {

	//Dekralasi array multidimensi
	var namaKaryawan = [2][4]string{{"Dina", "Bayu", "Dimas", "Robi"}, {"Riski", "Candini", "Bagas"}}
	// var nilai = [3][3]int{
	// 	{90, 98, 100},
	// 	{10, 28, 300},
	// 	{880, 99, 300},
	// }
	var dimensi3 = [2][2][2]int{
		{
			{
				1, 2,
			},
			{
				2, 2,
			},
		},
		{
			{
				3, 3,
			},
			{
				4, 4,
			},
		},
	}

	// fmt.Println(namaKaryawan)
	// fmt.Println(nilai)
	// fmt.Println(dimensi3)

	fmt.Println(namaKaryawan[0][2])
	fmt.Println(namaKaryawan[1][0])
	fmt.Println(dimensi3[1][0][1])
}
