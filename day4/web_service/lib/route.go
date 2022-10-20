package lib

import (
	"encoding/json"
	"net/http"
)

func Users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var hasil, err = json.Marshal(Data)

		if err == nil {
			w.Write(hasil)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Tidak bisa akses", http.StatusBadRequest)
	}
}

func User(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var username = r.FormValue("username")
	var hasil []byte
	var err error

	for _, d := range Data {
		if d.Username == username {
			hasil, err = json.Marshal(d)

			if err == nil {
				w.Write(hasil)
				return
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}
