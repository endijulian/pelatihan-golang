package main

import "fmt"

func main() {

	//Tipe data numerik (Desimal & non desimal)
	var nilai1 uint8 = 255
	nilai2 := 65535
	var nilai3 uint16 = 65535
	var nilai4 int8 = 127
	var nilai5 int = 93493489489448439

	fmt.Println(nilai1)
	fmt.Println(nilai2)
	fmt.Println(nilai3)
	fmt.Println(nilai4)
	fmt.Println(nilai5)

	//Desimal
	var desimal1 float32 = 67.86
	var desimal2 float64 = 5.8877
	desimal := 3.14

	fmt.Printf("Bilangan desimal : %f\n", desimal1)
	fmt.Printf("Bilangan desimal : %.2f\n", desimal2)
	fmt.Printf("Bilangan desimal : %.6f\n", desimal)

	//Tipe data bool
	var hujan bool = true
	cerah := false

	fmt.Printf("Apakah hari ini hujan ? %t \n", hujan)
	fmt.Printf("Apakah hari ini hujan ? %t \n", cerah)

	//String
	var text, pesan string
	text = "Selamat datang di rumah coding \n"
	pesan = `Belajar go lang
di Rumah Coding`

	fmt.Println(text)
	fmt.Println(pesan)
}
