package main

import (
	"fmt"
	"net/http"
)

//Function Untuk handler URL
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index Belajar Golang")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Selamat datang di Rumah Coding Endi Julian")
	})

	//Menerapkan function handling URL Custom
	http.HandleFunc("index/", index) //Paramater index adalah function buatan

	fmt.Println("Web server berjalan di http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
