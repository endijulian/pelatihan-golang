package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Nama":     "Joni Deep",
			"Kegiatan": "Belajar Golang",
		}

		var view, err = template.ParseFiles("./go.html")
		if err != nil {
			fmt.Printf("err.Error() : %v", err.Error())
		}

		view.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}
