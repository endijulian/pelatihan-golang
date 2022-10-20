package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	NamaPanjang string `json:"nama"`
	Alamat      string `json:"alamat"`
	Usia        int    `json:"usia"`
}

func main() {

	var jsonString = `{"nama": "Bayu Akbar", "usia": 30}`
	var jsonData = []byte(jsonString)

	//var struct data untuk menampung hasil Decode
	var data user

	err := json.Unmarshal(jsonData, &data)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Decode berhasil ditampung di struct data")
	fmt.Println(data)

	println()
	println()
	println()

	//Data Json Array
	var jsonArray = `[
		{"nama": "Bayu", "alamat": "Depok", "usia": 30},
		{"nama": "Akbar", "alamat": "Jakarta", "usia": 28}
	]`

	//Untuk menampung slice DECODE ARRAY ke dalam slice struct
	var dataSlice []user

	err1 := json.Unmarshal([]byte(jsonArray), &dataSlice)
	if err1 != nil {
		fmt.Printf("Error : %v\n", err1.Error())
	}

	fmt.Println("Decode Berhasil ditampung di slice data struct")
	fmt.Printf("Data Slice : %v\n", dataSlice)

	println()
	println()
	println()

	encodeJson, err2 := json.Marshal(dataSlice)
	if err2 != nil {
		fmt.Printf("Error : %v\n", err1.Error())
	}

	fmt.Println("Encode Berhasil")
	// fmt.Println("Hasil Encode: ", encodeJson)
	fmt.Println("Hasil Encode: ", string(encodeJson))

}
