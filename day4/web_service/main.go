package main

import (
	"fmt"
	"net/http"
	"web_service/lib"
)

func main() {
	http.HandleFunc("/users", lib.Users)
	http.HandleFunc("/user", lib.User)

	fmt.Println("Starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
