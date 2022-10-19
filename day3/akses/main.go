package main

import f "akses/library" // f berguna untuk alias / pengganti nama

func main() {

	//Menggunakan fungsi file lain
	f.Pesan("Endi")

	//Menggunakan struct dari file lain
	var kar1 = f.Karyawan{}
	kar1.Nama = "Joni"
	kar1.Alamat = "Depok"

}
