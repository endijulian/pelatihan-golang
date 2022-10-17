package main

import "fmt"

func main() {
	login := false

	//Penggunaan if
	if login == true {
		fmt.Println("Anda berhasil login")
	}

	//Penggunaan IF ELSE
	if login == true {
		fmt.Println("Anda berhasil login")
	} else {
		fmt.Println("Anda gagal login")
	}

	nilai := 90
	//IF ELSEIF
	if nilai < 80 {
		fmt.Println("Grade C")
	} else if nilai < 90 {
		fmt.Println("Grade B")
	} else if nilai <= 100 {
		fmt.Println("Grade A")
	} else {
		fmt.Println("Tidak ada predikat")
	}
}
