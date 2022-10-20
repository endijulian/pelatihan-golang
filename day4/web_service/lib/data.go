package lib

type Biodata struct {
	Username    string
	NamaPanjang string
	Usia        int
}

var Data = []Biodata{
	{Username: "User1", NamaPanjang: "Bayu Akbar", Usia: 30},
	{Username: "User2", NamaPanjang: "Dani Drajat", Usia: 70},
	{Username: "User3", NamaPanjang: "Oke Gesss", Usia: 20},
}
