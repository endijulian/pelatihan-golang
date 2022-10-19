package main

import (
	"fmt"
	"strings"
)

func main() {

	var dataKirim = []string{"John", "Wick", "Bayu", "Raysad"}
	var dataContainsO = filter(func(cekNilai string) bool { return strings.Contains(cekNilai, "a") }, dataKirim...)

	fmt.Println("Data Awal :", dataKirim)
	fmt.Println("Data Filter terdapat huruf a :", dataContainsO)

}

func filter(nilaiBalik func(string) bool, dataNilai ...string) (hasil []string) {
	for _, nilai := range dataNilai {
		filtered := nilaiBalik(nilai)
		if filtered == true {
			hasil = append(hasil, nilai)
		}
	}
	return
}
