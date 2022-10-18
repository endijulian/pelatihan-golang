package main

import "fmt"

func main() {
	var arrayMotor = [3]string{"Honda", "Yamaha", "Suzuki"}
	var sliceMotor = []string{"Honda", "Yamaha", "Suzuki"}

	copyArrayMotor := arrayMotor
	copySliceMotor := sliceMotor

	fmt.Println(copyArrayMotor)
	fmt.Println(copySliceMotor)

	//Ubah data slice
	// arrayMotor[0] = "Endi"
	// sliceMotor[0] = "Julian"

	copyArrayMotor[0] = "Endi"
	copySliceMotor[0] = "Julian"

	fmt.Println(arrayMotor)
	fmt.Println(sliceMotor)
}
