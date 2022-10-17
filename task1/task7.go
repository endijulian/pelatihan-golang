package main

import "fmt"

func main() {

	for i := 0; i < 51; i++ {
		fmt.Println("Angka :", i)
		if i == 5 {
			fmt.Println("Kelipatan 5 yang ke", i)
		}
	}

}
