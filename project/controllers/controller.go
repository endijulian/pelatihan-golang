package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"project/configs"
	"project/models"

	"github.com/gorilla/mux"
)

func GetAllStudents(w http.ResponseWriter, r *http.Request) {

	var students models.Students
	var arr_students []models.Students
	var response models.Response

	db := configs.Connect()
	defer db.Close()

	rows, err := db.Query("Select id, nama_depan, nama_belakang, no_hp, gender, jenjang, hobi, alamat from students")
	if err != nil {
		log.Print(err)
		return
	}

	for rows.Next() {
		if err := rows.Scan(&students.Id, &students.NamaDepan, &students.NamaBelakang, &students.NoHp, &students.Gender, &students.Hobi, &students.Alamat, &students.Jenjang); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_students = append(arr_students, students)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arr_students

	w.Header().Set("Content-Type", "application/json")
	//Mengubah data dari database menjadi json karena di encode
	json.NewEncoder(w).Encode(response)

}

func InsertStudent(w http.ResponseWriter, r *http.Request) {
	var response models.Response

	db := configs.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	nama_depan := r.FormValue("nama_depan")
	nama_belakang := r.FormValue("nama_belakang")
	no_hp := r.FormValue("no_hp")
	gender := r.FormValue("gender")
	jenjang := r.FormValue("jenjang")
	hobi := r.FormValue("hobi")
	alamat := r.FormValue("alamat")

	_, err = db.Exec("INSERT INTO students (nama_depan, nama_belakang, no_hp, gender, jenjang, hobi, alamat) values (?,?,?,?,?,?,?)", nama_depan, nama_belakang, no_hp, gender, jenjang, hobi, alamat)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Success Insert Data Siswa"
	log.Print("Berhasil Insert Data Siswa")

	w.Header().Set("Content-Type", "application/json")
	//Mengubah data dari database menjadi json karena di encode
	json.NewEncoder(w).Encode(response)
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	var students models.Students
	var response models.ResponseFind

	db := configs.Connect()
	defer db.Close()

	param := mux.Vars(r)
	id := param["id"]

	rows, err := db.Query("Select * From students Where id = ? ", id)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err := rows.Scan(&students.Id, &students.NamaDepan, &students.NamaBelakang, &students.NoHp, &students.Gender, &students.Hobi, &students.Alamat, &students.Jenjang)

		if err != nil {
			log.Print(err)
		}
	}

	response.Status = 1
	response.Message = "Success Get Data Siswa"
	response.Data = students
	log.Print("Berhasil Get Data Siswa")

	w.Header().Set("Content-Type", "application/json")
	//Mengubah data dari database menjadi json karena di encode
	json.NewEncoder(w).Encode(response)
}

func StudentUpdate(w http.ResponseWriter, r *http.Request) {
	// var response models.Response
	var students models.Students
	var response models.ResponseFind

	db := configs.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	param := mux.Vars(r)
	id := param["id"]
	// id := r.FormValue("id")
	nama_depan := r.FormValue("nama_depan")
	nama_belakang := r.FormValue("nama_belakang")
	no_hp := r.FormValue("no_hp")
	gender := r.FormValue("gender")
	jenjang := r.FormValue("jenjang")
	hobi := r.FormValue("hobi")
	alamat := r.FormValue("alamat")

	_, err = db.Exec("UPDATE students SET nama_depan = ?, nama_belakang  = ?, no_hp = ?, gender = ?, jenjang = ?, hobi = ?, alamat = ? WHERE id = ?", nama_depan, nama_belakang, no_hp, gender, jenjang, hobi, alamat, id)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Success Update Data Siswa"
	response.Data = students
	log.Print("Berhasil Update Data Siswa")

	w.Header().Set("Content-Type", "application/json")
	//Mengubah data dari database menjadi json karena di encode
	json.NewEncoder(w).Encode(response)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {

	var response models.ResponseDelete

	db := configs.Connect()
	defer db.Close()

	param := mux.Vars(r)
	id := param["id"]

	_, err := db.Exec("DELETE FROM students WHERE id = ?", id)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Success Delete Data Siswa"
	log.Print("Berhasil Delete Data Siswa")

	w.Header().Set("Content-Type", "application/json")
	//Mengubah data dari database menjadi json karena di encode
	json.NewEncoder(w).Encode(response)
}
