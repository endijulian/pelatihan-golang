package main

import (
	"fmt"
	"log"
	"net/http"
	c "project/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/students", c.GetAllStudents).Methods("GET")
	router.HandleFunc("/students/store", c.InsertStudent).Methods("POST")
	router.HandleFunc("/students/getId/{id}", c.GetStudents).Methods("GET")
	router.HandleFunc("/students/update/{id}", c.StudentUpdate).Methods("PUT")
	router.HandleFunc("/students/delete/{id}", c.DeleteStudent).Methods("DELETE")

	fmt.Println("Server Berjalan")
	log.Fatal(http.ListenAndServe(":8081", router))
}
