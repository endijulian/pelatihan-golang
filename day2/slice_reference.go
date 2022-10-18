package main

import (
	"fmt"
)

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

	//Penggunaan teknik 2 index [0:0]
	var mobil = []string{"Honda", "Mitshubishi", "Toyota"}   //Slice
	var mobil2 = [3]string{"Honda", "Mitshubishi", "Toyota"} //Array
	var mobilBaru = mobil[0:2]                               //Ambil index 0 sampai dengan sebelum index 2
	var mobilBaru2 = mobil2[0:2]                             //Ambil index 0 sampai dengan sebelum index 2

	fmt.Println("Nilai Slice Mobil", mobil)
	fmt.Println("Nilai Array Mobil 2", mobil2)

	mobilBaru[0] = "Daihatsu"
	mobilBaru2[0] = "Daihatsu"

	fmt.Println("Nilai Slice Mobil", mobil)
	fmt.Println("Nilai Array Mobil", mobil2)

}
