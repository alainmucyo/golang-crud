package database

import (
	"database/sql"
	"os"
)

var con *sql.DB

func Connect(DB_PASSWORD string, DB_USERNAME string, DB_DATABASE string, DB_TYPE string) *sql.DB {

	db, err := sql.Open(DB_TYPE, DB_USERNAME+":"+DB_PASSWORD+"@/"+DB_DATABASE)
	if err != nil {
		os.Exit(2)
	}
	con = db
	return db
}

func Conn() *sql.DB {
	return con
}
