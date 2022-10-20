package configs

import (
	"database/sql"
	"fmt"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/sekolah")

	if err != nil {
		fmt.Printf("Error :%v \n", err.Error())
	}

	return db
}
