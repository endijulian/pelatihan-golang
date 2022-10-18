package main

import "fmt"

func main() {
	var data = []map[string]string{
		{"nama": "Endi", "gender": "Laki-Laki"},
		{"nama": "Dwi", "gender": "Laki-Laki"},
		{"nama": "Agung", "gender": "Laki-Laki"},
		{"nama": "Jaka", "gender": "Laki-Laki"},
		{"nama": "Perdana", "gender": "Laki-Laki"},
		{"nama": "Ahmad", "gender": "Laki-Laki"},
		{"nama": "Rifandi", "gender": "Laki-Laki"},
		{"nama": "Ely", "gender": "Perempuan"},
		{"nama": "Handayani", "gender": "Perempuan"},
	}

	for key, value := range data {
		if key%2 == 1 {
			fmt.Printf("Index ke %v,\n Nama : %v\n Jenis Kelamin : %v\n", key, value["nama"], value["gender"])
		} else {
			// fmt.Println("Bilangan Genap", key)
		}
	}
}
