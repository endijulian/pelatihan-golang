package main

import "fmt"

func main() {

	hari := 2
	switch hari {
	case 1:
		fmt.Println("Senin")
		break
	case 2:
		fmt.Println("Selasa")
		break
	case 3:
		fmt.Println("Rabu")
		break
	case 4:
		fmt.Println("Kamis")
		break
	case 5:
		fmt.Println("Jumat")
		break
	case 6:
		fmt.Println("Sabtu")
		break
	case 7:
		fmt.Println("Minggu")
		break
	default:
		fmt.Println("Nan")
	}
}
