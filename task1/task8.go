package main

import "fmt"

func main() {

	// var start int = 10

	for a := 6; a > 0; a-- {

		for b := 6; b >= a; b-- {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}

	println()
	println()

	for brs := 1; brs <= 6; brs++ {

		for spasi := 6; spasi >= brs; spasi-- {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
